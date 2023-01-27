<template>
    <div style="min-height: 100px;margin: 0 30px;">
        <el-avatar :size="95" :src="this.$store.state.img"></el-avatar>
    </div>
   
    <el-container>
        <el-aside width="200px">
            <el-menu class="el-menu-vertical-demo">
                <el-menu-item index="1" @click="menuChange">
                    <span slot="title">修改资料</span>
                </el-menu-item>
                <el-menu-item index="2" @click="menuChange('UserCollection')">
                    <span slot="title">个人收藏</span>
                </el-menu-item>
                <el-menu-item index="3" @click="menuChange('ChangePassword')">
                    <span slot="title">修改密码</span>
                </el-menu-item>
            </el-menu>
        </el-aside>
        <el-main>
            <div>
                <component :is="mainComponent"></component>
            </div>
        </el-main>
    </el-container>



</template>
<script>
import UserCollection from './user_collection/UserCollection.vue';
import ChangePassword from './change_password/ChangePassword.vue';
export default {
    data() {
        return {
            mainComponent: "UserCollection",
           
        }
    },
    methods: {
       
        menuChange(mainComponent) {
            this.mainComponent = mainComponent
        },
    },
    components: {
        UserCollection,
        ChangePassword
    },
    name: 'AdminInformation'
}
</script>

<style scoped="scoped">

.el-menu-vertical-demo {
    background-color: rgba(214, 180, 69, 0.493);
    min-height: 750px;
}

.el-container {
    max-height: 750px;
}

.el-main {
    background-color: rgba(214, 180, 69, 0.493);
    max-height: 750px;
}

.el-aside {
    background-color: rgba(214, 180, 69, 0.493);
    max-height: 750px;
}
</style>