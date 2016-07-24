# gopm.in
The source code of gopm.in

## Goal
The goal of gopm.in is become the go version of pypi as python do, a central version control of go pkg and yet FASTER than gopkg.in

## URL Structure
As gopkg use 

"""
<username>/<app>.<version>
"""

as it's control method, which is good to understand and use.
However the problem is, it only support github. Therefor bitbucket and other VCS is not available.
It's important that we support those sites and not to increase the complexity of URL.
so one more layer of URL is a compromise solution.

"""
<gc/gh/bb/ja/ap/c>/<username>/<app>.<version>
"""

## Fast access
The following way to achive fast access all of the global:
1. CDN 
2. Static serving
