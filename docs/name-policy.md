# images name specification
- 拒绝使用 latest version images
```
    --nameRejectPolicy=latestTag 参数将默认拒绝 tag 为 latest 或为空
    忽略或不使用 name Policy: --nameRejectPolicy=""
```

- container 命名规范化进行警示, 不拒绝
```
    需要注意，当你的 ImagePolicyWebhook defaultAllow 设置为 false
    将会走默认策略，被拒绝
    可以通过不使用 latestTag 的 nameRejectPolicy 策略来避免
```
