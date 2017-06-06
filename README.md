# Zookeeper BOSH release

Uses BOSH [links](https://bosh.io/docs/links.html) as an example.

# Useful commands

```
$ bosh -d zookeeper run-errand smoke-tests
$ bosh -d zookeeper ssh -c '/var/vcap/jobs/zookeeper/bin/ctl status' -r
```
