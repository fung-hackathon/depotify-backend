# For Developers

## Coding

[Effective Go](https://golang.org/doc/effective_go.html)

commit 前に必ず `go fmt` でフォーマットをかける

## GitHub

### Branch Name

- `master`
  - リリース用。 stable。
- `dev`
  - 開発用ブランチ。latest。
- `feature`
  - 機能追加。devから分岐してdevにmergeする。`feature/{#issue}`
- `fix`
  - mainを修正するときに使用 `fix/{#issue}`

### Commit Message

```
{type} {summary}

{detail}
```

#### type
- `[add]`
  - 追加
- `[fix]`
  - 修正
- `[wip]`
  - 作業中
- `[clean]`
  - 整理
