# Go Vulnerability Scanner GitHub Action 🚀

Easily scan your Go projects for known vulnerabilities using the govulncheck tool provided by golang.org/x/vuln. This GitHub Action integrates seamlessly into your CI/CD pipeline, ensuring your dependencies are up-to-date and free from security risks.


## Usage
Add the following workflow file to your project:

.github/workflows/go-vuln-scan.yml:

```
name: Go Vulnerability Scan

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  vuln-scan:
    name: Run Go Vulnerability Scanner
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Code
        uses: actions/checkout@v4

      - name: Run Go Vulnerability Scanner
        uses: debug-ing/go-vuln-scanner@v1.0
```