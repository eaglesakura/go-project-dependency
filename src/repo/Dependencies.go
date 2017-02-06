package repo

import (
	"os"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"encoding/json"
)

type Dependencies struct {
	Repositories []Repository; // 依存しているリポジトリ一覧
}

// $GOPATHに存在するリポジトリを示すModel
type Repository struct {
	ImportPath string; // ImportしているライブラリPath "go get ${Import}"
	Rev        string; // Importしているライブラリのリビジョン "git checkout -f ${Revision}"
}

// gitのルートディレクトリである場合true
func getRepositoryHash(path string) string {
	cmd := exec.Command("git", "rev-parse", "HEAD");
	cmd.Dir = path;

	out, err := cmd.Output();
	if err != nil {
		return "";
	}
	result := string(out);
	result = result[:len(result) - 1];
	return result;
}

func checkoutRepositoryHash(path string, hash string) error {
	cmd := exec.Command("git", "checkout", "-f", hash);
	cmd.Dir = path;

	return cmd.Run();
}

func Find(values []os.FileInfo, action func(it os.FileInfo) (bool)) os.FileInfo {
	for _, it := range values {
		if action(it) {
			return it;
		}
	}
	return nil;
}

// リポジトリを検索する
func findRepository(path string, repo string, dep *Dependencies) {
	files, err := ioutil.ReadDir(path);
	if err != nil {
		return;
	}

	for _, info := range files {
		if (info.IsDir()) {
			fullPath := path + "/" + info.Name();
			child, _ := ioutil.ReadDir(fullPath);

			// リポジトリパスを足す
			var childRepo string;
			if len(repo) > 0 {
				childRepo = repo + "/" + info.Name();
			} else {
				childRepo = info.Name();
			}

			// ここがgit-rootであるかを確認する
			gitInfo := Find(child, func(it os.FileInfo) bool {
				return it.IsDir() && it.Name() == ".git";
			});


			if gitInfo == nil {
				// rootではないので、子を検索する
				findRepository(fullPath, childRepo, dep);
				continue;
			}

			// rootなので、ここで探索を終わる
			hash := getRepositoryHash(fullPath);
			if len(hash) != 0 {
				fmt.Printf("dependency repo[%s] hash[%s]\n", childRepo, hash);

				dumpRepo := Repository{ImportPath:childRepo, Rev:hash};

				// リポジトリを追加する
				dep.Repositories = append(dep.Repositories, dumpRepo);
			}
		}

	}
}


// 依存ライブラリを列挙する
func NewDependencies() (Dependencies, error) {
	GOPATH := os.Getenv("GOPATH");
	if len(GOPATH) == 0 {
		return Dependencies{}, errors.New("GOPATH not set");
	}

	fmt.Printf("GOPATH=%s\n", GOPATH);
	result := Dependencies{};
	findRepository(string(GOPATH) + "/src", "", &result);

	return result, nil;
}

func FromFile(path string) (Dependencies, error) {
	buf, err := ioutil.ReadFile(path);
	if err != nil {
		return Dependencies{}, errors.New(fmt.Sprint("FileError[%s]", path));
	}

	result := Dependencies{};
	if json.Unmarshal(buf, &result) != nil {
		return Dependencies{}, errors.New(fmt.Sprint("Json Unmershal Error[%s]", path));
	}
	return result, nil;
}

func (self *Dependencies) ToJson() string {
	buf, err := json.Marshal(self);
	if err != nil {
		return "";
	} else {
		return string(buf);
	}
}

func (self *Dependencies) ToFile(path string) error {
	buf, err := json.Marshal(self);
	if err != nil {
		return err;
	} else {
		return ioutil.WriteFile(path, buf, os.ModePerm);
	}
}

func (self *Dependencies) Restore() error {
	GOPATH := os.Getenv("GOPATH");
	if len(GOPATH) == 0 {
		return errors.New("GOPATH not set");
	}

	fmt.Printf("GOPATH=%s\n", GOPATH);

	// 全てのリポジトリをgetする
	for _, repo := range self.Repositories {
		fmt.Printf("go get %s\n", repo.ImportPath);
		cmd := exec.Command("go", "get", repo.ImportPath);
		cmd.Stdout = os.Stdout;
		cmd.Run();
	}

	// 全てのリポジトリをcheckoutする
	for _, repo := range self.Repositories {
		fmt.Printf("checkout %s [%s]\n", repo.ImportPath, repo.Rev);
		err := checkoutRepositoryHash(GOPATH + "/src/" + repo.ImportPath, repo.Rev);
		if err != nil {
			return err;
		}
	}

	// 全てのリポジトリを事前ビルドする
	for _, repo := range self.Repositories {
		fmt.Printf("go install %s\n", repo.ImportPath);
		cmd := exec.Command("go", "install", repo.ImportPath);
		cmd.Stdout = os.Stdout;
		cmd.Run();
	}

	return nil;
}