<template>
    <el-menu :default-active="activeIndex" mode="horizontal" class="el-menu-demo">
        <el-menu-item @click="getNovelList('玄幻')">玄幻</el-menu-item>
        <el-menu-item @click="getNovelList('武侠')">武侠</el-menu-item>
        <el-menu-item @click="getNovelList('网游')">网游</el-menu-item>
        <el-menu-item @click="getNovelList('穿越')">穿越</el-menu-item>
        <el-menu-item @click="getNovelList('科幻')">科幻</el-menu-item>
        <el-menu-item @click="getNovelList('悬疑')">悬疑</el-menu-item>
        <el-menu-item @click="getNovelList('其他')">其他</el-menu-item>
    </el-menu>
    <div class="card-bak" style="padding:0px 10px 20px 20px;">
        <el-row style="padding:0px 15px;">
            <el-col :span="4" v-for="(item) in novelList.slice((currentPage - 1) * PageSize, currentPage * PageSize)"
                @click="TitleInfor" :gutter="1">
                <div style="margin-top:40px;height: 200px; width: 200px;">
                    <el-card :body-style="{ padding: '2px', margin: '2px', height: '200px', width: '200px' }">
                        <div class="card-title">
                            <div>
                                <el-image :src="item.img" style="width: 100px;hight:100px;"></el-image>
                                <p style="width: 100px;hight:100px;" @click="AuthorInfor(item.author)">作者: {{
                                        item.author
                                }}</p>
                            </div>
                            <div class="right-title" >
                               <el-link @click="novelInformation(item.title)" style="width: 90px;hight:30px;font-weight: bold;font-size: 15px;" type="primary">{{ item.title }}</el-link>
                                <TextOver :value="item.introduce" style="width:90px;font-size: 4px;margin: 5px;" row="6"></TextOver>
                            </div>

                        </div>
                    </el-card>
                </div>
            </el-col>
        </el-row>
    </div>

    <div>
        <el-container class="page" style="float: right;">
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
import TextOver from "../TextOver.vue"

export default {
    name: 'NovelMain',
    data() {
        return {
            novelList: [],
            novelType: '',
            jumpPage: 1,
            total: 0,
            currentPage: 1,
            PageSize: 18,
        }
    },

    components: {
        TextOver,
    },
    created() {
        this.getNovelList(this.novelType)
    },
    methods: {
        async getNovelList(novelType) {
            const url = '/article/type?type=' + novelType
            const { data: res } = await this.$http.get(url)
            if (res.status !== 200) {
                if (res.status === 1004 || 1005 || 1006 || 1007) {

                }
                this.$message.error(res.message)
            }
            this.novelList = res.articles
            this.total = res.total

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

        AuthorInfor(author) {
            console.log(author)
        },
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
    }

}
</script>

<style scoped="scoped">




.card-title {
    display: flex;
    flex-wrap: wrap;
}

.card-bak {
    background-color: rgba(214, 180, 69, 0.493);

    height: 750px;
}


.right-title {
    justify-content: right;
    overflow: hidden;
    overflow-y: hidden;
}

.page {
    display: flex;
    justify-content: right;
}

.el-menu-demo {
    background-color: rgba(183, 185, 19, 0.493);
}
</style>