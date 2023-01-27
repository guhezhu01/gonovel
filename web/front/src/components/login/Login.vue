<template>
    <el-card style="height: 200px;">

    </el-card>

    <div class="login_form">
        <el-form :inline="true" :model="formInline" :rules="rules" ref="formInline" label-width="100px" style="margin:20px 400px ;justify-content: center;"
            class="demo-form-inline">
            <el-form-item label="用户名:" prop="username">
                <el-input v-model="formInline.username" autocomplete="off"></el-input>
            </el-form-item>
            <br>
            <el-form-item label="密码:" prop="password">
                <el-input v-model="formInline.password" autocomplete="off" show-password v-on:keyup.enter="onSubmit">

                </el-input>
            </el-form-item>
            <br>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">登录</el-button>
            </el-form-item>

            <el-form-item class="resetButton">
                <el-button type="danger" @click="resetForm('formInline')">取消</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
export default {
    data() {
        var validatePass = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入密码'))
            }
        }
        return {
            formInline: {
                username: '',
                password: '',
            },
            rules: {
                username: [{ required: true, message: "用户名不能为空！", trigger: 'blur' }],
                passWord: [{ validator: validatePass, trigger: 'blur' }],
            },

        }

    },
    methods: {
        async onSubmit() {
            //跳转到登录界面
            this.$refs.formInline.validate(async (valid) => {
                if (!valid) return this.$message.error('输入非法数据，请重新输入')
                const { data: res } = await this.$http({
                    method: "post",
                    url: "/login",
                    headers: {
                        'Content-Type': 'text/plain',
                    },
                    data: this.formInline,
                })
                if (res.status != 200) return this.$message.error(res.message)
                window.sessionStorage.setItem('token', res.token)
                this.$router.push('/type')
            })
        },

        resetForm(formIn) {
            this.$refs[formIn].resetFields()
        },

    },
    name: "Login"
}
</script>

<style>
.login_form{
    display: flex;
    align-items: center;
}
</style>