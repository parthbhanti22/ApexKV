# 🏔️ ApexKV

![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

**ApexKV** is a distributed, strongly consistent (CP) key-value store built from scratch in Go. 

It is engineered to demonstrate the core principles of distributed systems: **Consensus**, **Replication**, and **Fault Tolerance**. Unlike simple wrapper projects, ApexKV implements the **Raft Consensus Algorithm** and **LSM-Tree storage engine** without relying on heavy external libraries.

### ⚡ Key Features
* **Raft Consensus:** Leader election and log replication handling split-brain scenarios.
* **LSM-Tree Storage:** Disk-based persistence with O(1) write performance using an append-only log.
* **Custom Protocol:** A lightweight binary TCP wire protocol (no HTTP/JSON overhead).
* **Chaos Engineering:** Integrated "Chaos Monkey" testing suite to simulate node failures in real-time.
