<template>
    <el-container class="login">
      <el-aside width="200px">
        <el-image :src="src" 
          :fit="fit"
          class="login-image">        
        </el-image>
      </el-aside>
  
      <el-main>
        <el-form :inline="true" 
        :model="formInline" 
        :rules="rules"
        ref="formInline"
        label-width="100px"
        class="demo-form-inline">
          <el-form-item label="用户名:" prop="username">
            <el-input v-model="formInline.username"
            autocomplete="off"
            ></el-input>
          </el-form-item>
          <br>

          <el-form-item label="密码:" prop="password">
            <el-input v-model="formInline.password" 
            autocomplete="off"
            show-password
            v-on:keyup.enter="onSubmit" 
            >
            <!-- 按下enter执行onSubmit方法-->
            </el-input>
          </el-form-item>
          <br>
          <el-form-item>
            <el-button type="primary" @click="onSubmit" >登录</el-button>
          </el-form-item>

          <el-form-item  class="resetButton">
            <el-button type="danger" @click="resetForm('formInline')">取消</el-button>
            <transition name="el-fade-in-linear">
              <div v-show="show" class="transition-box">.el-fade-in-linear</div>
            </transition>
          </el-form-item>
        </el-form>
      </el-main>
    </el-container><br>
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
            username: [{ required:true,message:"用户名不能为空！", trigger: 'blur' }],
            passWord: [{ validator: validatePass,trigger: 'blur' }],
          },
          src: 'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg',
        }

      },
      methods: {
        async onSubmit() {
        this.$refs.formInline.validate(async (valid) => {
        if (!valid) return this.$message.error('输入非法数据，请重新输入')
        const { data: res } = await this.$http({
          method: "post",
          url:"/login",
          headers: {
                      'Content-Type':'text/plain',
                  },
          data:this.formInline,
        })
        if (res.status != 200) return this.$message.error(res.message)
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('/')})
        },
        resetForm(formIn) {
          this.$refs[formIn].resetFields()
        },
      }
    }
  </script>
  
<style scoped>
  .login{
    height: 100%;
  }
.el-aside {
        background-color: #d3dce6;
        color: var(--el-text-color-primary);
        text-align: center;
      }
  
.el-main {
      background-color: #e9eef3;
      color: var(--el-text-color-primary);
      text-align: center;
    }
.view{
  width: 100%;
  height: 4rem;
  line-height: 4rem;
  text-align: center;
}

</style>