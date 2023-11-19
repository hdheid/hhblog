# gvb_admin

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

## style
```css
*{
    padding: 0;
    margin:0;
    box-sizing: border-box;
}

这段CSS代码是一段通用的CSS样式，通常用于重置浏览器默认样式，以确保不同浏览器上的网页元素具有一致的外观和行为。让我解释这些样式的作用：

1. `*` 选择器：`*` 是通配符选择器，匹配文档中的所有元素。

2. `padding: 0;`：将元素的内边距（padding）设置为0。这将消除元素的默认内边距，确保元素的内部内容与元素边界之间没有空白。

3. `margin: 0;`：将元素的外边距（margin）设置为0。这将消除元素的默认外边距，确保元素之间没有不必要的间距。

4. `box-sizing: border-box;`：将元素的框模型（box model）设置为 "border-box"。这意味着元素的总宽度包括内边距和边框，而不是默认的 "content-box" 模型，其中内边距和边框会增加到元素的宽度之外。
                                                                                                                                          
 通过这些样式，你可以创建一个干净的起点，然后在网页中逐个添加自定义样式，以确保页面元素按照你的需求进行布局和显示。这是一种常见的CSS重置技巧，用于规范化不同浏览器之间的默认样式差异。
```

## html格式
```html
<!--admin的页面格式-->
<div class="gvb_admin">
    <aside>

    </aside>
    <div class="main">
        <header></header>
        <div class="tabs"></div>
        <main></main>
    </div>
</div>

<aside>元素通常用于包含侧边栏内容。
<header>元素通常用于包含页面的标题或页眉部分。
<div class="tabs">元素可能用于显示标签式导航选项卡。
<main>元素通常用于包含主要的页面内容。
```

## 格式化代码
快捷键：ctrl+alt+l

## 安装ui组件库
```shell
npm i --save ant-design-vue
```

### 在 vue 中引入该 ui 组件
在 `main.js` 中输入：
```js
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';

app.use(Antd)
```

### 图标的引入
去阿里巴巴的图标库里面：icfont。
将代码复制到 assets/css/iconfont.css 中，并在 main.js 中使用 `import "./assets/css/iconfont.css";`来引入，
即可通过 class 来进行使用 ，这里选取的是使用在线的方式。

### 字体图标的引入
官网：https://fontawesome.dashgame.com/
下载命令：`npm install font-awesome --save`
在 main.js 中使用 `import "font-awesome/css/font-awesome.min.css";` 来引入，
即可通过 class 来进行使用，这里选取的是下载到本地的方式。