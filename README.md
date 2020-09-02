你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
## AWS CloudFormation Resource Provider Go Plugin

The CloudFormation CLI (cfn) allows you to author your own resource providers that can be used by CloudFormation.

This plugin library helps to provide Go runtime bindings for the execution of your providers by CloudFormation.

Usage
-----

If you are using this package to build resource providers for CloudFormation, install the [CloudFormation CLI Go Plugin](https://github.com/aws-cloudformation/cloudformation-cli-go-plugin) - this will automatically install the the [CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli)! A Python virtual environment is recommended.

```bash
pip3 install cloudformation-cli-go-plugin
```

Refer to the documentation for the [CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli) for usage instructions.

Development
-----------

For changes to the plugin, a Python virtual environment is recommended. Check out and install the plugin in editable mode:

```bash
python3 -m venv env
source env/bin/activate
pip3 install -e /path/to/cloudformation-cli-go-plugin
```

You may also want to check out the [CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli) if you wish to make edits to that. In this case, installing them in one operation works well:

```bash
pip3 install \
  -e /path/to/cloudformation-cli \
  -e /path/to/cloudformation-cli-go-plugin
```

That ensures neither is accidentally installed from PyPI.

Linting and running unit tests is done via [pre-commit](https://pre-commit.com/), and so is performed automatically on commit. The continuous integration also runs these checks. Manual options are available so you don't have to commit:

```bash
# run all hooks on all files, mirrors what the CI runs
pre-commit run --all-files
# run unit tests only. can also be used for other hooks, e.g. black, isort, pytest-local
pre-commit run pytest-local
```

Getting started
---------------

This plugin create a sample Go project and requires golang 1.8 or above and [godep](https://golang.github.io/dep/docs/introduction.html). For more information on installing and setting up your Go environment, please visit the official [Golang site](https://golang.org/).


License
-------

This library is licensed under the Apache 2.0 License.
