this uses the goproxy modules to create a go module proxy.
the implementation, albeit incomplete, shall solve several issues:

1. keep a local cache of upstream software, so that we can guarantee repeatable builds

2. keep a local cache of upstream software, in case upstream software becomes unavailable

3. provide some authorisation workflow if new modules are added in an organisation so that they can be tracked, reviewed and audited when and if necessary

