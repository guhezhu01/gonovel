<template>
    <div>
        <div style="margin: 0px 50px;float: right;">
            <el-link :underline="false" @click="getRandNovel">
                <el-icon>
                    <Refresh />
                </el-icon> 换一批
            </el-link>
        </div>

        <div class="card-bak" style="padding:0px 10px 20px 20px;">
            <el-row style="padding:0px 15px;">
                <el-col :span="2" v-for="(item) in novelList" :offset="3" @click="TitleInfor" :align="top">
                    <div style="margin-top:40px;height: 200px; width: 200px;">
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
                                        style="width: 90px;hight:30px;font-weight: bold;font-size: 15px;"
                                        type="primary">{{ item.title }}</el-link>
                                    <TextOver :value="item.introduce" style="width:90px;font-size: 4px;margin: 5px;"
                                        row="6"></TextOver>
                                </div>

                            </div>
                        </el-card>
                    </div>
                </el-col>
            </el-row>
        </div>

    </div>
</template>
<script>
import TextOver from "../TextOver.vue"

export default {
    data() {
        return {
            novelList: [],


        }
    },
    components: {
        TextOver,
    },
    created() {
        this.getRandNovel()
    },
    methods: {
        async getRandNovel() {
            const { data: res } = await this.$http.get("article/rand")
            if (res.status === 200) {
                this.novelList = res.articles
            } else {
                this.$message.error("加载失败！")
            }
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
    },
    name: 'Home'

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
</style>