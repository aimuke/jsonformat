# jsonformat[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2699/badge)](https://bestpractices.coreinfrastructure.org/projects/2699)
这个工具用于进行json的格式化

主要是用于一些情况下需要比较两个json的时候，有时候两个json的字段顺序不一样，不便于进行比较。本工具主要是将输入的json先转化为对象，然后在将对象转化为字符串。由于使用相同的方式进行的转化，保证了json字段顺序的一致性。便于后续进行比较

## 生成可执行文件

- 在window中运行`build-linux.bat` 生成`format.sh`的可执行文件

- 在windows中运行 `build-win.bat` 生成`format.exe`可执行文件

## 使用方式
- 直接执行运行go代码

   在代码所在目录运行如下代码
   ```sh
    go run main.go -i a.json
   ```
- 可执行文件
   windows
  ```
  format.exe -i aa.json
  ```
  或 linux 中
  ```
  format.sh -i aa.json
  ```
   

## 参数
```sh
run.exe --help
Usage of format.exe:
  -0 string
        格式化后的文件位置, 默认 out.json (default "out.json")
  -i string
        需要格式化的文件位置, 默认 in.json (default "in.json")
  -indent
        输出是否进行格式化缩进, 默认为true (default true)
```
