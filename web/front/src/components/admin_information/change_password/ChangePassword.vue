<template>
    <el-form :inline="true" :rules="rules" ref="User" :model="User" label-width="100px" class="demo-form-inline">
        <el-form-item label="密码" prop="password">
            <el-input v-model="User.password" autocomplete="off" show-password></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="checkPassWord">
            <el-input v-model="User.checkPassWord" autocomplete="off" show-password></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="changePssword">确认修改</el-button>
        </el-form-item>

    </el-form>

</template>

<script>
export default {
    data() {
        var ISPWD = /^(?=.*[0-9])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*,\._\+(){}])[0-9a-zA-Z!@#$%^&*,\\._\+(){}]{8,20}$/;
        // 密码校验
        const validatePassword = (rule, value, callback) => {
            if (!ISPWD.test(this.User.password)) {
                callback(new Error("用户密码必须包含大写字母、小写字母、数字和特殊符号"));
            } else {
                callback();
            }
        }

        var validatePass2 = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请再次输入密码'))
            } else if (value !== this.User.password) {
                callback(new Error('两次输入密码不一致!'))
            } else {
                callback()
            }
        }
        return {
            User: {
                password: "",
                checkPassWord: "",
            },
            rules: {
                password: [
                    { required: true, trigger: "blur", message: "请输入您的密码" },
                    { min: 8, max: 20, message: '用户密码长度必须介于 8 和 20 之间', trigger: 'blur' },
                    { required: true, validator: validatePassword, trigger: 'blur' }
                ],
                checkPassWord: [{ validator: validatePass2, trigger: 'blur' }],
            }
        }

    },
    methods: {
        async changePssword() {
            this.$refs.User.validate(async (valid) => {
                if (!valid) return this.$message.error('输入非法数据，请重新输入')
                const { data: res } = await this.$http({
                    method: "post",
                    url: "/user/update/password",
                    headers: { 'Content-Type': 'application/json' },
                    data: {
                        "username": this.$store.state.userInformation.username,
                        "password": this.User.password,
                        "id": this.$store.state.userInformation.ID
                    }
                })
                console.log(res)
            })
        }
    }

}
</script>

<style>

</style>