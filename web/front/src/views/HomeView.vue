<template>
  <div class="bg" :style="bg">
    <div class="home">
      <el-container>
        <el-header>
          <Header></Header>
        </el-header>

        <el-container>
          <el-main>
            <router-view />
          </el-main>
        </el-container>
      </el-container>
    </div>
  </div>
</template>

<script>
import Header from '@/components/Header.vue';
export default {
  data() {
    return {
      bg: {
        backgroundImage: "url(" + require("../assets/123.jpeg") + ")",
        backgroundRepeat: "no-repeat",
        backgroundSize: "100% 100%",
      }
    }
  },
  name: 'HomeView',
  components: {
    Header,

  }
}
</script>

<style>
.el-header {
  background-color: #89cd227a;
  color: var(--el-text-color-primary);
  text-align: center;
  height: 60px;
}

.el-aside {
  background-color: #d3dce6;
  color: var(--el-text-color-primary);
  text-align: center;
  height: 900px;
}

.el-main {
  background-color: #e9eef3;
  color: var(--el-text-color-primary);
  height: 900px;

}
</style>