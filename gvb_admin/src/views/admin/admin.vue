<template>
    <div class="gvb_admin">
        <aside>
            <div>

            </div>
        </aside>
        <div class="main">
            <header>
                <div class="left">
                    <a-breadcrumb>
                        <a-breadcrumb-item>首页</a-breadcrumb-item>
                        <a-breadcrumb-item><a href="#">个人中心</a></a-breadcrumb-item>
                        <a-breadcrumb-item><a href="#">用户列表</a></a-breadcrumb-item>
                        <a-breadcrumb-item>更多</a-breadcrumb-item>
                    </a-breadcrumb>
                </div>
                <div class="right">
                    <div class="icon_actions">
                        <i class="fa fa-home"></i>
                        <i v-if="theme" class="fa fa-sun-o" @click="setTheme"></i>
                        <i v-else class="fa fa-moon-o" @click="setTheme"></i>
                        <i class="fa fa-arrows-alt"></i>
                    </div>
                    <div class="avatar">
                        <img src="https://gss0.baidu.com/9vo3dSag_xI4khGko9WTAnF6hhy/zhidao/pic/item/a2cc7cd98d1001e900e51e21bb0e7bec55e797c6.jpg"
                             alt="头像">
                    </div>
                    <div class="drop_menu">
                        <a-dropdown placement="bottom">
                            <!--加了个placement="bottom"可以让列表居中，所以要将header的padding设置为40px，否则会变样子-->
                            <a class="ant-dropdown-link" @click.prevent>
                                菜单
                                <i class="fa fa-bars"></i>   <!--这里并没有鼠标移动上去变成手的特效，但是在antd的官网上有（下拉菜单）-->
                            </a>
                            <template #overlay>
                                <a-menu @click="menuClick"> <!-- 下拉菜单点击事件 -->
                                    <a-menu-item key="user_center">
                                        <i class="fa fa-user-circle-o"></i>
                                        <a href="javascript:;" style="margin-left: 5px;">个人中心</a>
                                        <!--5px 可以让两个标签之间来点空间-->
                                    </a-menu-item>
                                    <a-menu-item key="2">
                                        <i class="fa fa-commenting"></i>
                                        <a href="javascript:;" style="margin-left: 5px;">我的消息</a>
                                    </a-menu-item>
                                    <a-menu-item key="article_list">
                                        <i class="fa fa-file-text"></i>
                                        <a href="javascript:;" style="margin-left: 5px;">文章列表</a>
                                    </a-menu-item>
                                    <a-menu-item key="logout">
                                        <i class="fa fa-sign-out"></i>
                                        <a href="javascript:;" style="margin-left: 5px;">注销退出</a>
                                    </a-menu-item>
                                </a-menu>
                            </template>
                        </a-dropdown>
                    </div>
                </div>
            </header>
            <div class="tabs"></div>
            <main></main>
        </div>
    </div>
</template>

<script setup>
import {useRouter} from "vue-router";
import {ref} from "vue";
//实例化一个useRouter
const router = useRouter()

function menuClick({key}) {
    if (key === "logout") {
        alert("调用注销接口！")
        return
    }
    router.push({
        name: key
    })
}

const theme = ref(true) //true表示白天，false表示晚上

function setTheme() {
    theme.value = !theme.value

    if (theme.value) {
        // 白天
        document.documentElement.classList.remove("dark") //当为白天的时候，删掉dark，如果本来没有也不会报错
    } else {
        document.documentElement.classList.add("dark") //将class设置为dark
    }
}

</script>

<style lang="scss">
.gvb_admin {
    width: 100%;
    display: flex;

    aside {
        width: 240px;
        height: 100vh;
        background-color: #2b3539
    }

    .main {
        width: calc(100% - 240px);

        header {
            height: 60px;
            background-color: white;
            padding: 0 40px;
            display: flex;
            justify-content: space-between;
            align-items: center;

            .right {
                display: flex;
                align-items: center;
                //margin-right: 20px; //可以让右边的图标再往中间收一点
            }

            .icon_actions {
                margin-right: 20px;

                i {
                    margin-left: 10px; //图标之间的间距
                    cursor: pointer; //这个表示当光标移动到该元素上时候，光标变成手，表示该元素可点击
                    font-size: 20px;
                    color: var(--text)
                }

                i:hover { //这边暂时用不了，后续问一下,（问题解决了，是因为 app.vue 里面设置常量的时候没有加分号）
                    color: var(--active);
                }
            }

            .avatar {
                img {
                    width: 40px;
                    height: 40px;
                    border-radius: 50%;
                }
            }

            .drop_menu {
                margin-left: 10px; //设置菜单和头像之间的间隔
            }
        }

        .tabs {
            height: 30px;
            border: 1px solid #f0eeee;
        }

        main {
            background-color: var(--bg); //在theme.css里面
            height: calc(100vh - 90px);
        }
    }
}

</style>