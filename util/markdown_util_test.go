package util

import "testing"

func TestConvertLuteMarkdown(t *testing.T) {
	markdown := "## Guide\n" +
		"\n" +
		"这是一篇讲解如何正确使用 **Markdown** 的排版示例，学会这个很有必要，能让你的文章有更佳清晰的排版。\n" +
		"\n" +
		"> 引用文本：Markdown is a text formatting syntax inspired\n" +
		"\n" +
		"## 语法指导\n" +
		"\n" +
		"### 普通内容\n" +
		"\n" +
		"这段内容展示了在内容里面一些排版格式，比如：\n" +
		"\n" +
		"- **加粗** - `**加粗**`\n" +
		"- *倾斜* - `*倾斜*`\n" +
		"- ~~删除线~~ - `~~删除线~~`\n" +
		"- `Code 标记` - `` `Code 标记` ``\n" +
		"- [超级链接](https://hacpai.com) - `[超级链接](https://hacpai.com)`\n" +
		"- [username@gmail.com](mailto:username@gmail.com) - `[username@gmail.com](mailto:username@gmail.com)`\n" +
		"\n" +
		"### 提及用户\n" +
		"\n" +
		"@Vanessa 通过 `@User` 可以在内容中提及用户，被提及的用户将会收到系统通知。\n" +
		"\n" +
		"> NOTE:\n" +
		">\n" +
		"> 1. @用户名之后需要有一个空格\n" +
		"> 2. 新手没有艾特的功能权限\n" +
		"\n" +
		"### 表情符号 Emoji\n" +
		"\n" +
		"支持大部分标准的表情符号，可使用输入法直接输入，也可手动输入字符格式。通过输入 `:` 触发自动完成，可在个人设置中[设置常用表情](https://hacpai.com/settings/function)。\n" +
		"\n" +
		"#### 一些表情例子\n" +
		"\n" +
		":smile: :laughing: :dizzy_face: :sob: :cold_sweat: :sweat_smile:  :cry: :triumph: :heart_eyes: :relieved:\n" +
		":+1: :-1: :100: :clap: :bell: :gift: :question: :bomb: :heart: :coffee: :cyclone: :bow: :kiss: :pray: :anger:\n" +
		"\n" +
		"### 大标题 - Heading 3\n" +
		"\n" +
		"你可以选择使用 H1 至 H6，使用 ##(N) 打头。建议帖子或回帖中的顶级标题使用 Heading 3，不要使用 1 或 2，因为 1 是系统站点级，2 是帖子标题级。\n" +
		"\n" +
		"> NOTE: 别忘了 # 后面需要有空格！\n" +
		"\n" +
		"#### Heading 4\n" +
		"\n" +
		"##### Heading 5\n" +
		"\n" +
		"###### Heading 6\n" +
		"\n" +
		"### 图片\n" +
		"\n" +
		"```\n" +
		"![alt 文本](http://image-path.png)\n" +
		"![alt 文本](http://image-path.png \"图片 Title 值\")\n" +
		"```\n" +
		"\n" +
		"支持复制粘贴直接上传。\n" +
		"\n" +
		"### 代码块\n" +
		"\n" +
		"#### 普通\n" +
		"\n" +
		"```\n" +
		"*emphasize*    **strong**\n" +
		"_emphasize_    __strong__\n" +
		"var a = 1\n" +
		"```\n" +
		"\n" +
		"#### 语法高亮支持\n" +
		"\n" +
		"如果在 ``` 后面跟随语言名称，可以有语法高亮的效果哦，比如:\n" +
		"\n" +
		"##### 演示 Go 代码高亮\n" +
		"\n" +
		"```go\n" +
		"package main\n" +
		"\n" +
		"import \"fmt\"\n" +
		"\n" +
		"func main() {\n" +
		"\tfmt.Println(\"Hello, 世界\")\n" +
		"}\n" +
		"```\n" +
		"\n" +
		"##### 演示 Java 高亮\n" +
		"\n" +
		"```java\n" +
		"public class HelloWorld {\n" +
		"\n" +
		"    public static void main(String[] args) {\n" +
		"        System.out.println(\"Hello World!\");\n" +
		"    }\n" +
		"\n" +
		"}\n" +
		"```\n" +
		"\n" +
		"> Tip: 语言名称支持下面这些: `ruby`, `python`, `js`, `html`, `erb`, `css`, `coffee`, `bash`, `json`, `yml`, `xml` ...\n" +
		"\n" +
		"### 有序、无序、任务列表\n" +
		"\n" +
		"#### 无序列表\n" +
		"\n" +
		"- Java\n" +
		"  - Spring\n" +
		"    - IoC\n" +
		"    - AOP\n" +
		"- Go\n" +
		"  - gofmt\n" +
		"  - Wide\n" +
		"- Node.js\n" +
		"  - Koa\n" +
		"  - Express\n" +
		"\n" +
		"#### 有序列表\n" +
		"\n" +
		"1. Node.js\n" +
		"   1.1. Express\n" +
		"   1.2. Koa\n" +
		"   1.3. Sails\n" +
		"2. Go\n" +
		"   2.1. gofmt\n" +
		"   2.2. Wide\n" +
		"3. Java\n" +
		"   3.1. Latke\n" +
		"   3.2. IDEA\n" +
		"\n" +
		"#### 任务列表\n" +
		"\n" +
		"- [x] 发布 Sym\n" +
		"- [X] 发布 Solo\n" +
		"- [ ] 预约牙医\n" +
		"\n" +
		"### 表格\n" +
		"\n" +
		"如果需要展示数据什么的，可以选择使用表格。\n" +
		"\n" +
		"| header 1 | header 3 |\n" +
		"| -------- | -------- |\n" +
		"| cell 1   | cell 2   |\n" +
		"| cell 3   | cell 4   |\n" +
		"| cell 5   | cell 6   |\n" +
		"\n" +
		"### 隐藏细节\n" +
		"\n" +
		"<details>\n" +
		"<summary>这里是摘要部分。</summary>\n" +
		"这里是细节部分。\n" +
		"</details>\n" +
		"\n" +
		"### 段落\n" +
		"\n" +
		"空行可以将内容进行分段，便于阅读。（这是第一段）\n" +
		"\n" +
		"使用空行在 Markdown 排版中相当重要。（这是第二段）\n" +
		"\n" +
		"### 数学公式\n" +
		"\n" +
		"$$a^2 + b^2 = \\color{red}c^2$$\n" +
		"\n" +
		"### 流程图\n" +
		"\n" +
		"```mermaid\n" +
		"graph TB\n" +
		"    c1-->a2\n" +
		"    subgraph one\n" +
		"    a1-->a2\n" +
		"    end\n" +
		"    subgraph two\n" +
		"    b1-->b2\n" +
		"    end\n" +
		"    subgraph three\n" +
		"    c1-->c2\n" +
		"    end\n" +
		"```\n" +
		"\n" +
		"### 时序图\n" +
		"\n" +
		"```mermaid\n" +
		"sequenceDiagram\n" +
		"    Alice->>John: Hello John, how are you?\n" +
		"    loop Every minute\n" +
		"        John-->>Alice: Great!\n" +
		"    end\n" +
		"```\n" +
		"\n" +
		"### 甘特图\n" +
		"\n" +
		"```mermaid\n" +
		"gantt\n" +
		"    title A Gantt Diagram\n" +
		"    dateFormat  YYYY-MM-DD\n" +
		"    section Section\n" +
		"    A task           :a1, 2019-01-01, 30d\n" +
		"    Another task     :after a1  , 20d\n" +
		"    section Another\n" +
		"    Task in sec      :2019-01-12  , 12d\n" +
		"    another task      : 24d\n" +
		"```\n" +
		"\n" +
		"### 图表\n" +
		"\n" +
		"```echarts\n" +
		"{\n" +
		"  \"title\": { \"text\": \"最近 30 天\" },\n" +
		"  \"tooltip\": { \"trigger\": \"axis\", \"axisPointer\": { \"lineStyle\": { \"width\": 0 } } },\n" +
		"  \"legend\": { \"data\": [\"帖子\", \"用户\", \"回帖\"] },\n" +
		"  \"xAxis\": [{\n" +
		"      \"type\": \"category\",\n" +
		"      \"boundaryGap\": false,\n" +
		"      \"data\": [\"2019-05-08\",\"2019-05-09\",\"2019-05-10\",\"2019-05-11\",\"2019-05-12\",\"2019-05-13\",\"2019-05-14\",\"2019-05-15\",\"2019-05-16\",\"2019-05-17\",\"2019-05-18\",\"2019-05-19\",\"2019-05-20\",\"2019-05-21\",\"2019-05-22\",\"2019-05-23\",\"2019-05-24\",\"2019-05-25\",\"2019-05-26\",\"2019-05-27\",\"2019-05-28\",\"2019-05-29\",\"2019-05-30\",\"2019-05-31\",\"2019-06-01\",\"2019-06-02\",\"2019-06-03\",\"2019-06-04\",\"2019-06-05\",\"2019-06-06\",\"2019-06-07\"],\n" +
		"      \"axisTick\": { \"show\": false },\n" +
		"      \"axisLine\": { \"show\": false }\n" +
		"  }],\n" +
		"  \"yAxis\": [{ \"type\": \"value\", \"axisTick\": { \"show\": false }, \"axisLine\": { \"show\": false }, \"splitLine\": { \"lineStyle\": { \"color\": \"rgba(0, 0, 0, .38)\", \"type\": \"dashed\" } } }],\n" +
		"  \"series\": [\n" +
		"    {\n" +
		"      \"name\": \"帖子\", \"type\": \"line\", \"smooth\": true, \"itemStyle\": { \"color\": \"#d23f31\" }, \"areaStyle\": { \"normal\": {} }, \"z\": 3,\n" +
		"      \"data\": [\"18\",\"14\",\"22\",\"9\",\"7\",\"18\",\"10\",\"12\",\"13\",\"16\",\"6\",\"9\",\"15\",\"15\",\"12\",\"15\",\"8\",\"14\",\"9\",\"10\",\"29\",\"22\",\"14\",\"22\",\"9\",\"10\",\"15\",\"9\",\"9\",\"15\",\"0\"]\n" +
		"    },\n" +
		"    {\n" +
		"      \"name\": \"用户\", \"type\": \"line\", \"smooth\": true, \"itemStyle\": { \"color\": \"#f1e05a\" }, \"areaStyle\": { \"normal\": {} }, \"z\": 2,\n" +
		"      \"data\": [\"31\",\"33\",\"30\",\"23\",\"16\",\"29\",\"23\",\"37\",\"41\",\"29\",\"16\",\"13\",\"39\",\"23\",\"38\",\"136\",\"89\",\"35\",\"22\",\"50\",\"57\",\"47\",\"36\",\"59\",\"14\",\"23\",\"46\",\"44\",\"51\",\"43\",\"0\"]\n" +
		"    },\n" +
		"    {\n" +
		"      \"name\": \"回帖\", \"type\": \"line\", \"smooth\": true, \"itemStyle\": { \"color\": \"#4285f4\" }, \"areaStyle\": { \"normal\": {} }, \"z\": 1,\n" +
		"      \"data\": [\"35\",\"42\",\"73\",\"15\",\"43\",\"58\",\"55\",\"35\",\"46\",\"87\",\"36\",\"15\",\"44\",\"76\",\"130\",\"73\",\"50\",\"20\",\"21\",\"54\",\"48\",\"73\",\"60\",\"89\",\"26\",\"27\",\"70\",\"63\",\"55\",\"37\",\"0\"]\n" +
		"    }\n" +
		"  ]\n" +
		"}\n" +
		"```\n" +
		"\n" +
		"### 五线谱\n" +
		"\n" +
		"```abc\n" +
		"X: 24\n" +
		"T: Clouds Thicken\n" +
		"C: Paul Rosen\n" +
		"S: Copyright 2005, Paul Rosen\n" +
		"M: 6/8\n" +
		"L: 1/8\n" +
		"Q: 3/8=116\n" +
		"R: Creepy Jig\n" +
		"K: Em\n" +
		"|:\"Em\"EEE E2G|\"C7\"_B2A G2F|\"Em\"EEE E2G|\\\n" +
		"\"C7\"_B2A \"B7\"=B3|\"Em\"EEE E2G|\n" +
		"\"C7\"_B2A G2F|\"Em\"GFE \"D (Bm7)\"F2D|\\\n" +
		"1\"Em\"E3-E3:|2\"Em\"E3-E2B|:\"Em\"e2e gfe|\n" +
		"\"G\"g2ab3|\"Em\"gfeg2e|\"D\"fedB2A|\"Em\"e2e gfe|\\\n" +
		"\"G\"g2ab3|\"Em\"gfe\"D\"f2d|\"Em\"e3-e3:|\n" +
		"```\n" +
		"\n" +
		"## 快捷键\n" +
		"\n" +
		"我们的编辑器支持很多快捷键，具体请参考 [键盘快捷键](https://hacpai.com/article/1474030007391)（或者按 \"`?` \":smirk_cat:）"
	t.Log(ConvertLuteMarkdown(markdown))
}
