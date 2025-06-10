# Pull Request Checklist

Please ensure you have completed the following:

- [ ] No `.env` variable has been pushed.
- [ ] All `fmt.Print()` or `log.Print()` debug statements are removed.
- [ ] No sensitive information (API keys, passwords, etc.) is present in the code or config.
- [ ] All new code is covered by unit tests where applicable.
- [ ] `go fmt` and `go vet` have been run on all Go files.
- [ ] No unused imports or variables remain.
- [ ] All error returns are properly handled.
