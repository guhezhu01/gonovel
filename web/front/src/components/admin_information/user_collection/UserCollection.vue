<template>
    <div class="card-bak" style="padding:40px 10px 20px 20px;margin: 40px;">
        <el-row style="padding:0px 15px;">
            <el-col :span="6" v-for="(item) in useCollInf.slice((currentPage - 1) * PageSize, currentPage * PageSize)"
                @click="TitleInfor" :gutter="2">
                <div style="margin-top:40px;height: 200px; width: 200px;" @mouseenter="optCollection()" @mouseleave="">
                    <el-card :body-style="{ padding: '2px', margin: '2px', height: '200px', width: '200px' }">
                        <div class="card-title">
                            <div>
                                <el-image :src="item.img" style="width: 100px;hight:100px;"></el-image>
                                <p style="width: 100px;hight:100px;" @click="AuthorInfor(item.author)">作者: {{
                                    item.author
                                }}</p>
                            </div>
                            <div class="right-title">
                                <el-link @click="novelInformation(item.title)"
                                    style="width: 90px;hight:30px;font-weight: bold;font-size: 15px;" type="primary">{{
                                        item.title
                                    }}</el-link>
                                <TextOver :value="item.introduce" style="width:90px;font-size: 4px;margin: 5px;"
                                    row="6"></TextOver>

                                <el-affix position="bottom">
                                    <el-button type="warning" size="mini" style="float: right;"
                                        @click="deleteCollection(item)">取消收藏</el-button>
                                </el-affix>
                            </div>
                        </div>
                    </el-card>
                </div>
            </el-col>
        </el-row>
    </div>
    <div v-if="dialogVisible">
        <el-container class="page" style="margin: 20px 50px;">
            <el-button type="primary" @click="jumpFirstPage" style="float: left;">首页</el-button>
            <el-pagination background prev-text="上一页" next-text="下一页" layout="total,prev, pager, next,jumper"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" :total="total"
                :current-page="currentPage" :page-size="PageSize">
            </el-pagination>
            <el-button type="primary" @click="jumpLastPage" style="float: right;margin: 2px;">尾页</el-button>
        </el-container>
    </div>
</template>

<script>
import TextOver from "../../TextOver.vue"
import {VueCropper} from 'vue-cropper'
export default {
    data() {
        return {
            useCollInf: [],
            jumpPage: 1,
            total: 0,
            currentPage: 1,
            PageSize: 8,
            dialogVisible: false,
            colleDialogVisible: false
        }
    },
    created() {
        this.getCollection()
    },
    components: {
        TextOver,
        VueCropper
    },
    methods: {
        async novelInformation(title) {
            const { data: res } = await this.$http(
                {
                    method: "post",
                    url: "/article",
                    headers: { 'Content-Type': 'application/json' },
                    data: {
                        title: title,
                    }
                }
            )
            this.$store.state.novelInformation = res.article
            const url = "/novel_information?title=" + title
            this.$router.push(url)
        },
        async deleteCollection(item) {


            const { data: res } = await this.$http({
                method: "delete",
                url: "/user/collection/delete?user_id=" + this.$store.state.userInformation.ID +
                    "&article_id=" + item.id + "&type=" + item.type,
                headers: { 'Content-Type': 'multipart/form-data' },

            })
            if (res.status == 200) {
                this.useCollInf = this.useCollInf.filter(i => i.id != item.id || i.type != item.type);
                this.total -= 1
                this.$message.success("已取消收藏")
            } else {
                this.$message.warn(res.message)
            }
        },
        async getCollection() {
            const { data: res } = await this.$http({
                methods: "get",
                url: "/user/collection/?user_id=" + this.$store.state.userInformation.ID,
            })
            if (res.status == 200) {
                this.useCollInf = res.data
                this.total = res.total
                if (this.total > 8) {
                    this.dialogVisible = true
                }
            } else {
                this.$message.error(res.message)
            }

        },
        jumpFirstPage() {
            this.currentPage = 1;
        },
        jumpLastPage() {
            this.currentPage = Math.ceil(this.total / this.PageSize);
        },
        // 显示第几页
        handleCurrentChange(val) {
            // 改变默认的页数
            this.currentPage = val
        },
        optCollection() {
            this.colleDialogVisible = true
        }
    },
    name: "UserCollection"

}
</script>

<style>
.card-title {
    display: flex;
    flex-wrap: wrap;
}
</style>