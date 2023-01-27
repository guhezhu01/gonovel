<template>
  <el-menu active-text-color="#ffd04b" background-color="#545c64" class="el-menu-vertical-demo" default-active="2"
    text-color="#fff" @open="handleOpen" @close="handleClose" router>
    <el-sub-menu index="1">
      <template #title style="font-size: 10px;">
        <el-icon>
          <Edit />
        </el-icon>
        <span>文章管理</span>
      </template>
      <el-menu-item-group>
        <el-menu-item index="/article">
          <el-icon>
            <EditPen />
          </el-icon>文章编辑
        </el-menu-item>
        <el-menu-item index="1-2">
          <el-icon>
            <Operation />
          </el-icon>文章审核
        </el-menu-item>
      </el-menu-item-group>

      <el-sub-menu index="1-4">
        <template #title>item four</template>
        <el-menu-item index="1-4-1">item one</el-menu-item>
      </el-sub-menu>
    </el-sub-menu>
    <el-menu-item index="/comment">
      <el-icon><ChatLineSquare /></el-icon>
      <span>评论管理</span>
    </el-menu-item>
    <el-menu-item index="/user">
      <el-icon>
        <Avatar />
      </el-icon>
      <span>用户管理</span>
    </el-menu-item>
    <el-menu-item index="/self">
      <el-icon>
        <Setting />
      </el-icon>
      <span>个人设置</span>
    </el-menu-item>

  </el-menu>

</template>
  
<script>
export default {
  data() {
    return {
      isCollapse: true
    }
  },
  methods: {

  },
  name: "AsideMenu",
}
</script>

<style scoped>
.el-menu-vertical-demo {

  height: 100%;
}
</style>