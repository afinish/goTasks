# Git

## Basic commands

#### git init
	examples: git init
Description: Initializes new repository in the current directory.

#### git init [project name]
	example: git init [project name]
Description: Initializes new repo inside the specified new directory.

#### git clone
	examples:
	git clone username@host:/path/to/repository
	git clone /path/to/repository
Description: *git clone* gets a copy of a repo inside the specified path
In the first example, the command copies a repo from a remote server, in the latter 
it copies it from a local directory.

#### git add
	example: git add <temp.txt>
Description: adds the specified file to the staging space. The example command adds the 'temp.txt' to the staging area.

#### git commit
	examplegit commit -m "Message to go with the commit"
Description: creates a commit with the changes and saves them in the git directory.
	
#### git status
	example: git status
Description: shows the status of the files (e.g. whether they are in the staging area or not, differences between the committed files and the current file, etc.)

#### git push
	example: git push origin <master>
Description: pushes/sends the changes/commits to the git repository (you can specify the branch). The command in the example sends the commits to the master branch of the git repository.

#### git checkout
	example: git checkout -b <branch-name>
Description: used for creating branches and navigating between them. The command in the example creates a new branch and switches to it.

#### git remote
	examples: 
	git remote -v
	git remote add origin <host-or-remoteURL>
	git remote rm <name-of-the-repository>
Description: used to connect the local repo with the repo in the specified address (example 2), delete the connection (example 3), and list all connections (example 1).

#### git branch
	examples:
	git branch
	git branch <branch-name>
	git branch -d <branch-name>
Description: used to create (example 2), delete (example 3), and list branches (example 1).

#### git pull
	example: git pull
Description: used the pull/merge the changes from the remote git repo to the local directory.

#### git merge 
	example: git merge <branch-name>
Description: used to merge the working branch with the specified branch.

#### git diff 
	examples:
	git diff --base <file-name>
	git diff <source-branch> <target-branch>
	git diff
Description: used to display the differences/conflicts between the base file and the specified file (example 1), between branches (example 2), list all the differences/conflicts (example 3).

#### git tag
	example: git tag <insert-commitID-here>
Description: mostly used to tag/mark the commits with the major releases (e.g. V1.0, V2.0)

#### git log
	example: git log
Description: used to display commit history of the repository

#### git reset
	example: git reset --hard HEAD
Description: used to reset the current directory to the last commit's state

#### git rm
	example: git rm filename.txt
Description: deletes the specified file

#### git stash
	example: git stash
Description: temporarily saves the changes to the local directory. It is done to save the progress to come back to it later without making commits.

#### git show
	example: git show
Description: shows the specified git objects (commits, tags, trees, blobs).

#### git fetch
	example: git fetch origin
Description: used to fetch the objects from the remote repo which are not present in the current repository. ???

#### git ls-tree 
	example: git ls-tree HEAD
Description: ???

#### git cat-file
	example: git cat-file -p d670460b4b4aece5915caf5c68d12f560a9fe3e4
Description: used to display metadata about the specified file or repository object.

#### git grep
	example: git grep "www.google.com"
Description: used to find the specified information in the files inside the committed trees, working directory, and staging area.

#### gitk
	example: gitk
Description: used to display the git GUI for local repository.

#### git instaweb
	example: git isntaweb -httpd-webrick
Description: used to display the git web interface local repository.

#### git gc
	example: git gc
Description: used to clean unnecessary files and optimize the local repository.

#### git archive
	example: git archive --format=tar master
Description: used to create .zip or .tar archives of the specified repository tree.

#### git prune
	example: git prune
Description: used to delete objects without incoming pointers.
Reference: example: 
Related: *git fetch*

#### git fsck
	example: git fsck
Description: used to check whether the git file system has any corrupted files

#### git rebase
	example: git rebase master
Description: sometimes used instead of merge. It works by 'changing the history'. Basically, is commits the changes in the secondary branch to the master branch without merging.
