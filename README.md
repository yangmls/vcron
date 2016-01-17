VCron library
================================================

It is still in development, do not use it now

# What is VCron

VCron is a replacement of linux cron, it is designed for distributed system and build with golang.

# Features

* Simple and no database dependency

* Use protobuf to send and receive data, it is much smaller than json

* Use json to store config asynchronously

* High performance by golang such as goroutine

* Agent can be ran without server so it is ok to disconnect by accident

# How to install

Run on server

	go get github.com/yangmls/vcron-server
	vcron-server --port=7023 --panel=8080

Run on agent

	go get github.com/yangmls/vcron-agent
	vcron-agent --name=agent_name --port=7023

# TODO

* Sync lock

* Panel pages

* Heartbeat and reconnect when losing

* Agent can run cron independently

* even more
