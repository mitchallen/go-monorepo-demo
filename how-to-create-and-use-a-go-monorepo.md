# How to Create and Use a Go Monorepo (Golang, Workspaces)

In this article, I will show how to create and use a Go monorepo.

<aside>
💡 These instructions were written and tested on a Mac.

</aside>

## Step 1. Setup requirements

For this article, I am using:

- Visual Studio Code (VS Code)
- Docker
- Dev Container Extension in VS Code
- Go version 1.21 or higher

## Step 2. Create a project folder

- Create a project folder:

```bash
mkdir -p ~/projects/golang/go-monorepo-demo
cd ~/projects/golang/go-monorepo-demo
```

## Step 3. Start Docker

- Start Docker (`open -a Docker`)
- Open VS Code (note dot)

```bash
code .
```

## Step 4. Add a .gitignore file

For this step, you will need a **.gitignore extension**, like the one found here:

- [https://marketplace.visualstudio.com/items?itemName=codezombiech.gitignore](https://marketplace.visualstudio.com/items?itemName=codezombiech.gitignore)

Once that extension is installed, do the following:

- Add a Go-specific **.gitignore** file
- In VS Code, call up the palette with this key combo: **Cmd-Shift-P**
- Type in **Add gitignore** and select it
- Type in **Go** and select it

This will generate a Go-specific **.gitignore** file.

## Step 5. Add Dev Container

- If you don’t have the dev container extension installed - see my article referenced previously
- Open the command palette (**Cmd-Shift-P**)
- Search for: **Dev Containers: Add Dev Container Configuration Files**
- Select **Go** as the container configuration template
- If not listed, select **Show All**
- Select the default version (must be 1.20 or higher)
- Click **OK** for any other default options
- When it appears, click **Reopen in Container**
- Open up the terminal window in VS Code (**Ctrl-`** or **View / Terminal**)

<aside>
💡 To run inside the dev container: you should execute all commands inside the VS Code terminal window. But if you run into issues, like accessing github, etc. – use an external terminal window for those commands.

</aside>

## Step 6. Init the workspace

In recent versions of Go, workspaces were introduced, including the **go work** command.  So make sure you are running a recent version for this step.

- Run this command:

```sh
go work init
```

- That command will create a new **go.work** file in the root of your project
- To view the new file run this command:

```sh
cat go.work
```

## Step 7. Create a package

*In the steps below you will need to change any reference to my git host, username, and project name to yours before running some commands:*

- Make a new package subfolder:

```bash
mkdir -p ./pkg/alpha
```

- To initialize the package, change this command to match your git provider, git provider username, and project folder name:

```bash
go mod init -C ./pkg/alpha github.com/mitchallen/go-monorepo-demo/pkg/alpha
```

<aside>
💡 Notice the use of the recently introduced **-C** flag. You can now run some commands in subfolders, without having to switch to that folder. It’s very handy for working in monorepos.

</aside>

- Create a source and test file for the package:

```bash
touch ./pkg/alpha/alpha.go
touch ./pkg/alpha/alpha_test.go
```

<aside>
💡 Nothing says that you have to use **/pkg** as the subfolder name - or that you need to use a subfolder at all. That’s just my preference for individual modules in a monorepo. You can use **/lib** or whatever you want. There’s also a special use case for folders called **/internal** (links at the end of this article).

</aside>

### Define alpha.go

- Paste this code into **/pkg/alpha/alpha.go** and save it:

```go
// Author: Mitch Allen
// File: alpha.go

package alpha

import (
    "fmt"

    "github.com/mitchallen/coin"
)

func Hello() {
    fmt.Println("[alpha]: Hello!")
}

func CoinCount(limit int) map[bool]int {

    m := map[bool]int{
        true:  0,
        false: 0,
    }

    for i := 0; i < limit; i++ {
        m[coin.Flip()]++
    }

    return m
}
```

What the code itself does isn’t that important.  The point was just to have something that exports a function.  To make it interesting I also added a dependency for one of my go packages.  You don’t need to install it.  I’ll show you later in this step how to get it to install automatically.

### Define alpha_test.go

- Paste this code into **/pkg/alpha/alpha_test.go** and save it:

```go
/**
 * Author: Mitch Allen
 * File: alpha_test.go
 */

package alpha

import (
    "math"
    "testing"
)

func TestCountCount(t *testing.T) {

    limit := 100
    expectedThreshold := int(math.Round(float64(limit) * 0.3))

    if got := CoinCount(limit); got[true] < expectedThreshold {
        t.Errorf("CoinCount(%d) = %v (threshold: %d)", limit, got, expectedThreshold)
    }
}
```

### Update the workspace

- Update the workspace (note the dot at the end):

```bash
go work use -r .
```

- To see the update that it generated for **go.work**, run this command:

```bash
cat go.work
```

### Clean things up and install the external package

Whenever I want to make sure things are working, I run **go mod tidy**.  You’ll find many times that external dependencies get cached or out of date. It’s a good way to clean things up.  It will also go and get missing dependencies.

- Run **go mod tidy** to download dependencies, etc.

```sh
go mod tidy -C ./pkg/alpha/
```

- It should find and download the referenced external package

<aside>
💡 When publishing changes to your own dependencies, you will find that you need to run **go mod tidy** in your consuming apps. Otherwise, your code may not be using the latest tags. If all else fails, **go get** with the specific tag.

</aside>

## Step 8. Make a scripts folder

With a monorepo, simple things like running tests get more complicated.  To simplify things I’m going to show you how to wrap complex calls into a script file. The script file can be bundled with your monorepo.

- Run the following commands to create a script folder and test script:

```sh
mkdir scripts
touch scripts/test.sh
chmod +x scripts/test.sh
```

- Paste the contents below into **test.sh** and save it:

```sh
#!/bin/bash

go list -f '{{.Dir}}' -m | xargs -L1 go mod tidy -C
go list -f '{{.Dir}}' -m | xargs -L1 go work sync -C
go list -f '{{.Dir}}' -m | xargs -L1 go test -C
```

The first half of each command lists the path of each module and returns the path on its own line (you can run that part on it’s own to see what it produces).

The results are then piped into the second command (**xargs**) which uses the **-C** flag (introduced in recent versions of **go**) to pass the full path from the list command to one of three go commands (**go mod tidy**, **go work sync**, and **go test**). The last one runs the tests.

- Run the test script:

```sh
 ./scripts/test.sh
```

Hopefully, the **tidy** command will be successful and all the tests will pass.

## Step 9. Create a git repo

One reason to use a monorepo is to keep all dependencies in one project.  You don’t have to export anything for use.  But if you would like to use a monorepo to make a bundle of modules available, you need to do and consider the following:

- Publish your git-based monorepo on github or wherever you prefer (gitlab, bitbucket, etc)
- Create and push tags for all changes to your monorepo (I’ll show you how in the next step)
- The important thing is that you can create a repo that matches your package declaration
- I used **github.com/mitchallen/go-monorepo-demo** for this article example, because I knew that was where I was going to post the code
- If another program tries to go get your **alpha** package, it will use your path to try to find it

<aside>
💡 It can be tricky having the proper access to update github from within a Dev Container - so maybe init and publish the repo and run the next tagging step outside of VS Code.

</aside>

## Step 10. Tag the build

Once you’ve created a git repo for your project, you should tag it to give it a version number.

- To do that, run these commands:

```sh
git tag v0.1.0
git push origin --tags
```

Every time you publish new changes, run that again, updating the version number (0.1.0, 0.1.1, 0.1.2, etc.)

## Step 11. Create a test project

If your monorepo is self-contained, you may never need to worry about testing the modules (packages) as a dependency for another project.  But if you would like to, do the following:

- Create a second Go project:

```sh
mkdir -p ~/projects/golang/go-monorepo-test
cd ~/projects/golang/go-monorepo-test
```

You can keep it simple.  We just need to create a test project to use our monorepo. If you don’t have an up-to-date version of **go** installed, consider using a Dev Container.

- Initialize the module (update the command for your git user, host and project name):

```sh
go mod init github.com/mitchallen/go-monorepo-test
```

- Create a **main.go** file:

```sh
touch main.go
```

- Paste the code below into **main.go** and save it
- You can either edit the reference to the ***alpha*** package to the location of your monorepo – or just leave it pointing to mine if you have nothing set up:

```go
package main

import (
    "fmt"

    "github.com/mitchallen/go-monorepo-demo/pkg/alpha"
)

func main() {
    fmt.Println(alpha.CoinCount(200))

    alpha.Hello()
}
```

- Run **go mod tidy** on the new repo

Hopefully **tidy** installed the latest dependency and you didn’t run into version issues.

- Run it:

```sh
go run main.go
```

## Step 12. Create additional modules

A monorepo isn’t much of a monorepo if it only has one module. So for illustrative purposes, add these additional modules:

- **pkg/beta** - export at least one function
- **cmd/service1** - a command line program using **main**
- **cmd/service2** - another command line program

Just remember that every time you add a new module you will need to update **go.work** with this command (note the dot at the end):

```sh
go work use -r .
```

<aside>
💡 Just like with **/pkg**, nothing says you have to use **/cmd** for a command subfolder name. Again, that’s just my preference. Which I may even change at some point.

</aside>

For **beta**, use the same pattern as for **alpha**. Make up whatever new function you would like.

For the two command modules, you can just make simple versions of Hello World.

The functionality is not that important.  The goal is just to show a combination of modules living in the same monorepo.

### Dependencies in a monorepo

You should make at least one of the command modules depend on **alpha** or **beta**.  In fact, you could make **beta** also depend on **alpha**.

Grouping modules together that are interdependent is one of the main reasons developers use monorepos.  For example, if you need to go make changes to **alpha**, and retest them in **beta**, it’s easier to rerun things together as one project as you work out the bugs.  Your tests will also catch problems immediately.  No need to publish, tag, tidy, and repeat simultaneously in two different projects.

<aside>
💡 If you don’t intend to use a module outside of a monorepo you won’t need to init it with a path (go mod init host/user/path/pkg/module). But to future-proof it, it’s probably a good idea. However not doing it may be one way to hide the module from the outside world.

</aside>

### Tag your changes

If you want to make **beta** or any other new module available externally, push a new tag. Then test it using the new module in your test program. Remember that you may need to use **go mod tidy** (good to run it anyway) and run **go get** manually using the new tagged version to get your test app to use the latest version.

## Cloning tips

If you look at the **.gitignore** file you will notice that it filters out the **go.work** file.

When someone clones your monorepo, they will need to rebuild it with this command:

```bash
go work init
go work use -r .
```

They should also run your test script to make sure everything works locally:

```bash
./scripts/test.sh
```

Consider making those steps part of the usage section in your project’s README file.

## Example Repo

- You can find the example repos here:
- [https://github.com/mitchallen/go-monorepo-demo](https://github.com/mitchallen/go-monorepo-demo) - monorepo
- [https://github.com/mitchallen/go-monorepo-test](https://github.com/mitchallen/go-monorepo-test) - optional test consumer

## Conclusion

In this article, you learned how to:

- Use Go workspaces to initialize a monorepo
- Create new modules for a Go monorepo
- Run tests and other functions against all the modules in a Go monorepo
- Setup a scripts folder to handle complex tasks
- Create a second project to use the monorepo as a dependency
- Add more modules, dependencies and commands to make it a true monorepo

## References

- [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces
- [Go: Project structure — MonoRepo](https://blog.devops.dev/go-project-structure-monorepo-daa762ec36a2)
- [Why I use the internal folder for a Go-project](https://medium.com/@as27/internal-folder-133a4867733c)
- [Don’t Put All Your Code In Internal](https://ido50.net/content/dont-put-all-your-code-in-internal)
- [cmd/go: provide a convenient way to iterate all packages in go.work workspace](https://github.com/golang/go/issues/50745)
