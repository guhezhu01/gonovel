<template>
    <el-form ref="form" :model="form" label-width="150px">
        <el-form-item label="用户名">
            <el-input style="width:20%;" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="QQ:">
            <el-input style="width:20%;" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="个人简介">
            <el-input type="textarea" style="width:25%;" maxlength="50" show-word-limit v-model="form.desc" rows="3">
            </el-input>
        </el-form-item>

        <el-form-item>
            <div class="block" default="头像" >
                <el-avatar :size="100" :src="circleUrl" />
                <el-button type="primary" style="text-align: center;" @click="">上传</el-button>
            </div>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="onSubmit">更新</el-button>
        </el-form-item>
    </el-form>
</template>
  
<script>
export default {
    data() {
        return {
            form: {
                name: '',
                region: '',
                date1: '',
                date2: '',
                delivery: false,
                type: [],
                resource: '',
                desc: '',
            },
        }
    },
    methods: {
        onSubmit() {
            console.log('submit!')
        },
    },
}
</script>