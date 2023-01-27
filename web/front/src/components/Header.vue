<template>
    <div class="header">
        <el-dialog title="欢迎登录！" v-model="addDialogVisible" width="30%" destroy-on-close center>
            <el-form :inline="true" :rules="rules" ref="user" :model="user" label-width="100px"
                class="demo-form-inline">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="user.username" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="user.password" autocomplete="off" show-password></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="LoginSubmit">确 定</el-button>
                    <el-button @click="addDialogVisible = cancelBut(user)">取 消</el-button>
                </span>
            </template>
        </el-dialog>

        <el-dialog title="欢迎注册！" v-model="registerDialogVisible" width="30%" destroy-on-close center>
            <el-form :inline="true" :rules="rules" ref="user" :model="user" label-width="100px"
                class="demo-form-inline">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="user.username" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="user.password" autocomplete="off" show-password></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="LoginSubmit">确 定</el-button>
                    <el-button @click="addDialogVisible = cancelBut(user)">取 消</el-button>
                </span>
            </template>
        </el-dialog>


        <h1 class="title" :default-active="1">guhezhu</h1>
        <el-menu class="el-tab" mode="horizontal" router>
            <el-menu-item style="font-size: 25px;" index="/">首页</el-menu-item>
            <el-menu-item style="font-size: 25px;" index="/type">分类</el-menu-item>
        </el-menu>
        <el-input class="search-title" placeholder="请输入" v-model="searchTitle"></el-input>
        <el-button class="but-title" @click="searchFunc">搜索</el-button>
        <el-button v-if="this.$store.state.showname" style="float: right;margin:10px;" @click="Register">注册</el-button>
        <el-button v-if="this.$store.state.showname" style="float: right;margin:10px;" @click="Login">登录</el-button>
        <el-button type="danger" v-if="!this.$store.state.showname" style="float: right;margin:10px;" @click="Logout">
            退出登录</el-button>
        <el-avatar style="float: right;margin:10px;" :size="size" :src="this.$store.state.img">
        </el-avatar>
        <el-button type="primary" v-if="!this.$store.state.showname" style="float: right;margin:10px;"
            @click="goAdminInformation">
            个人中心</el-button>
    </div>

</template>

<script>

export default {
    data() {
        return {
            searchTitle: '',
            addDialogVisible: false,
            registerDialogVisible: false,
            user: {
                username: '',
                password: '',
            },

            rules: {
                username: [
                    { required: true, trigger: "blur", message: "请输入您的账号" },
                    { min: 2, max: 20, message: '用户账号长度必须介于 2 和 20 之间', trigger: 'blur' }
                ],
                password: [
                    { required: true, trigger: "blur", message: "请输入您的密码" },
                    { min: 8, max: 20, message: '用户密码长度必须介于 8 和 20 之间', trigger: 'blur' },
                    { required: true, trigger: 'blur' }
                ],

            },
        }
    },
    methods: {
        goAdminInformation() {
            this.$router.push('/admin')
        },
        async Register() {
            this.registerDialogVisible = true
        },
        async searchFunc() {
            if (this.searchTitle === "") {
                this.$message.error("请输入内容！")
            } else {
                this.searchTitle = this.searchTitle.replace(/\s/g, "")
                const url = "/novel_information?title=" + this.searchTitle
                const { data: res } = await this.$http(
                    {
                        method: "post",
                        url: "/article",
                        headers: { 'Content-Type': 'application/json' },
                        data: {
                            title: this.searchTitle,
                        }
                    }
                )
                if (res.status === 200) {
                    this.$store.state.novelInformation = res.article
                    this.$router.push(url)
                } else {
                    this.$message.error("小说不存在！")
                }

            }
        },
        Login() {
            this.addDialogVisible = true
        },
        async LoginSubmit() {
            this.$refs.user.validate(async (valid) => {
                if (!valid) return this.$message.error('输入非法数据，请重新输入')
                const { data: res } = await this.$http({
                    method: "post",
                    url: "/login",
                    headers: {
                        'Content-Type': 'text/plain',
                    },
                    data: this.user,
                })
                if (res.status != 200) return this.$message.error(res.message)
                window.sessionStorage.setItem('token', res.token)
                this.$store.state.userInformation = res.userInformation
                this.$store.state.showname = false

                const { data: res01 } = await this.$http({
                    method: "get",
                    url: "/user/information/?userId=" + this.$store.state.userInformation.ID,

                })
                this.$store.state.img = res01.information.img
                this.$message.success("登录成功！")
                this.addDialogVisible = false


            })
        },
        cancelBut(userTable) {
            Object.keys(userTable).forEach((key) => (userTable[key] = ''))
            return false
        },
        Logout() {
            localStorage.clear()
            this.$store.state.img = ""
            this.$store.state.showname = true
            this.$store.state.userInformation = {}
            this.user.password = ''
        }


    },
    name: "Header"
}
</script>

<style scoped="scoped">
.title {
    color: rgba(222, 82, 164, 0.772);
    float: left;
}

.search-title {
    margin: 10px;
    width: 15%;
    height: 40px;
    text-align: center;
}

.but-title {
    width: 3%;
    height: 40px;
}

.header {
    max-height: 60px;
}

.el-tab {
    width: 15%;
    height: 40px;
    float: left;
    margin: 10px;
    background-color: #c7d15e7a;
}
</style>