# Questions and Answers

If your question is not covered here, please send it to us [on GitHub](https://github.com/noxworld-dev/noxscript/discussions/new?category=q-a)
or in our [Discord](https://discord.gg/HgDUeXhAyW).

## How do I start?

For a quick guide on using NoxScript, see [quickstart](./quickstart.md).

## What's the difference with existing NoxScript 3?

See [this guide](./ns3_to_go.md).

## Where can I find the examples?

See [examples](../examples).

## How does it all work?

As you may know OpenNox supports original NoxScript, which is compiled into the map.
Nothing changed here: it still runs NoxScript from the map even if new scripts are available.

Additionally, OpenNox includes a Go language interpreter, which scans map folder for Go files.
OpenNox also bundles all supported script runtimes and dependencies. Thus, there's no need to have any Go tooling
to run these scripts. Scripts are just text files distributed with the map.

Usually, when player joins a Nox server, a compressed `.map` file is sent (`.nxz`).
OpenNox extends this process: instead of sending map the regular way, it runs a tiny web server for map downloads.

If client is also OpenNox, server makes a `.zip` archive of the map folder and sends it (filtering out unrelated files).
Thus, all Go scripts will travel with the map in this archive.

And if the client is a good-old vanilla Nox, OpenNox server sends map as usual. In this case, unfortunately,
all Go scripts are lost for the client. The scripts still run on the server, so client will still see their effect.
It's just that the client won't be able to host these map without OpenNox (and re-downloading the map).