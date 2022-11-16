# Starter User API

This api provides user management capabilities. It exposes endpoints for:
- Create new user
- Change password
- Delete user
- Give access to user
- Log in user

### Runnning Local
- Run `make run_local`

### API Documentation
This repo uses [gin-swagger](https://github.com/swaggo/gin-swagger), you need to have swagger installed, so plese run:
- `make install_swagger`
- To generate swagger documentation run: `make doc_swagger`

In order to generate documentation for your api add comments to your handlesr in the format described [here](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).

## !IMPORTANT: For contributors
If you are a contributor, please follow the steps below to enable `git hooks` used by this project:

- Run `make config_git_hooks` to set the `git hooks` folder to [project's git hook folder](./.githooks). 

### To commit: 
This project expects that the commit message needs to comply with semantic version commit message conventions. So it is expected that your commit messages follows the conventions described in [Semantic Version Tools Repo](https://github.com/GUILN/semver).

To do so, [commit-msg hook](./.githooks/commit-msg) will enforce this convention.

**It is impotant** to follow this convention as we use it to generate the version of the package based on `semantic version` standards: `v.major.minor.patch` according to the `semantic version commit message`, this enables any commit to be eligible for generating a new version of the package.

### To generate and publish versions of the package:
- Make sure you are in `main branch` 
- Make sure your last commit message describes and indicates the expected semantics for the version to be generated.
- Run `make new_version`

### **!NOTE: If it is the first time you are pushing to this repo**.
You need to create the first version:
- Run `git tag -a v0.0.1 -m "first version"`
- Run `git push origin --tags v0.0.1`

**This steps will make this version available for the users.**
