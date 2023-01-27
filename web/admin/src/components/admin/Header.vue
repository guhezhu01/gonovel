<template>
    <p>后台管理</p>
    <el-button type="danger" @click="logout" class="a" plain>退出</el-button>
</template>
  
  <script>
  export default {
    methods: {
        //退出清除token
      logout() {
        window.sessionStorage.clear('token')
        this.$router.push('login')
      }
    },
    name:"Header"
  }
  </script>
  
  
<style scoped>
.a{
    position: absolute;
    right:10px;
    top:20px;
}
</style>