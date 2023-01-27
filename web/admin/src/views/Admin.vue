<template>
    <el-container class="home_con">
      <el-header>
        <div><Header></Header></div>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <div><AsideMenu></AsideMenu></div>
        </el-aside>
        <el-main>
          <div> <router-view /></div>
        </el-main>
      </el-container>
    </el-container>
  </template>
  
  <script>
  import AsideMenu from '../components/admin/AsideMenu.vue'
  import Header from '../components/admin/Header.vue'
    export default {
      data(){
        return {
        }
      },
      components: {
        AsideMenu,
        Header,
    },
    }
  </script>
  
  <style scoped="scoped">
  .home_con{
    height: 100%;
  }
  .el-header {
      background-color: #b3c0d1;
      color: var(--el-text-color-primary);
      text-align: center;
      line-height: 60px;
    }
  
  .el-aside {
      background-color: #545c64;
      color: var(--el-text-color-primary);
      text-align: center;
      display: flex;
      flex-direction: column;
    }
  
  .el-main {
      background-color: #e9eef3;
      color: var(--el-text-color-primary);
      text-align: center;
     
    }
  .item{
   
    width: 200px;
    background-color: powderblue;
  }
  
  .el-footer {
      background-color: #b3c0d1;
      color: var(--el-text-color-primary);
      text-align: center;
    }
  
  </style>