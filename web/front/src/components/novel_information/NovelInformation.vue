<template>

    <el-breadcrumb separator="/" style="font-size: 20px;">
        <el-breadcrumb-item @click="goback">
            <el-link style="font-size: 20px;">返回</el-link>
        </el-breadcrumb-item>
        <el-breadcrumb-item>小说信息</el-breadcrumb-item>
    </el-breadcrumb>
    <div class="novel_infor">

        <el-image :src="novelInformation.img" style="width: 200px;hight:300px;padding:30px 50px; "></el-image>
        <div class="infor">
            <div style="display: inline-flex;">
                <p style="font-size: 25px;font-weight: bold;">{{ novelInformation.title }}</p>
                <p style="font-size: 25px; ">/{{ novelInformation.author }}</p>
            </div>
            <el-row :gutter="6">
                <el-col :span="2">
                    <el-tag type="success" size="large">{{ novelInformation.type }}</el-tag>
                </el-col>
                <el-col :span="3">
                    <el-tag type="success" size="large">{{ novelInformation.clicks }}</el-tag>
                </el-col>
                <el-col :span="3">
                    <el-tag type="success" size="large">{{ novelInformation.chapter }}</el-tag>
                </el-col>
                <el-col :span="3">
                    <el-tag type="success" size="large">{{ novelInformation.words }}</el-tag>
                </el-col>
                <el-col :span="4">
                    <el-tag type="success" size="large">{{ novelInformation.state }}</el-tag>
                </el-col>
            </el-row>
            <div style="max-height: 80px;">
                <p>{{ novelInformation.introduce }}</p>
            </div>
            <div style="margin: 90px 0px;">
                <el-button type="warning" @click="collectionNovel(novelInformation.id, novelInformation.type)">
                    加入收藏<el-icon>
                        <Star />
                    </el-icon>
                </el-button>

                <el-button type="primary" @click="downloadNovel(novelInformation.id, novelInformation.title)">
                    下载<el-icon>
                        <Download />
                    </el-icon>
                </el-button>

            </div>
        </div>
        <el-divider style="margin: 0px 0px;"></el-divider>

        <el-card class="comment" style="width: 100%;min-height: 420px;margin: 0px 0px;text-align: 0px 0px;">
            <div style="display:flex ;"> <el-avatar :src=this.$store.state.img style="margin :10px 30px"></el-avatar>
                <el-input style="width: 50%;" type="textarea" :rows="3" placeholder="请输入评论" maxlength="100"
                    v-model="userComment">
                </el-input>
                <el-button type="primary" style="margin:10px 10px ;" @click="addComment(userComment)">添加评论</el-button>
            </div>

            <div style="display: flex;margin :20px  30px" v-for="comment in comments">
                <el-row :gutter="200">
                    <el-col :span="1">
                        <el-avatar :src=comment.img>
                        </el-avatar>
                    </el-col>
                    <el-col :span="1" style="padding: 0 0;">
                        <div style="display: flex;font-size: 10px;">
                            <span> {{ comment.username }}</span>
                        </div>

                        <div style="display: flex;word-break:keep-all;white-space:nowrap;">
                            <div>{{ comment.content }}</div>
                            <div style="padding: 0px 20px;">
                                <i class="element-icons el-icon-icon" @click="addAgree(comment)"
                                    :disabled="codeDisabled" style="font-size: 20px;"><span style="font-size: 10px;">{{
                                        comment.agrees
                                    }}</span></i>
                            </div>
                        </div>
                    </el-col>

                </el-row>
            </div>
        </el-card>

    </div>
</template>
<script>
import { Download } from '@element-plus/icons-vue'


export default {
    data() {
        return {
            count: 0,
            novelInformation: [],
            userComment: "",
            jumpPage: 1,
            total: 0,
            currentPage: 1,
            PageSize: 18,
            comments: [],
        }
    },
    created() {
        this.getNovelInformation()
        this.getComment()
    },
    methods: {
        async getComment() {
            const { data: res } = await this.$http({
                method: "get",
                url: "/comments?article_id=" + this.$store.state.novelInformation.id + "&title=" + this.$store.state.novelInformation.title,
            })
            if (res.status == 200) {
                this.comments = res.comments
            }

        },
        async addComment(userComment) {
            if (userComment !== "") {
                if (this.$store.state.userInformation.ID != null) {

                    const { data: res } = await this.$http(
                        {
                            method: "post",
                            url: "/comment/add",
                            headers: { 'Content-Type': 'application/json' },
                            data: {
                                user_id: this.$store.state.userInformation.ID,
                                article_id: this.$store.state.novelInformation.id,
                                article_title: this.$store.state.novelInformation.title,
                                username: this.$store.state.userInformation.username,
                                content: userComment,
                            }
                        }
                    )
                    if (res.status === 200) {
                        this.$message.success("评论成功!")
                        this.userComment = ""
                        this.getComment()
                    } else {
                        this.$message.error(res.message)
                    }
                } else {
                    this.$message.error("请登录!")
                }
            } else {
                this.$message.warning("请输入内容！")
            }
        },

        async addAgree(comment) {
            comment.agrees++
            const { data: res } = await this.$http({
                method: "post",
                url: "/comment/add/agrees/" + comment.id + "/" + comment.user_id,
                headers: { 'Content-Type': 'application/json' },
                data: {
                    "agrees": comment.agrees,
                }
            })


        },

        getNovelInformation() {
            this.novelInformation = this.$store.state.novelInformation
        },
        goback() {
            this.$store.state.novelInformation = []
            this.$router.go(-1)
        },
        async downloadNovel(id, title) {
            const url = "/article/download?id=" + id + "&title=" + title
            this.$http.get(url, {
                responseType: "arraybuffer"
            }).then((response) => {

                const blob = new Blob([response.data])
                const blobUrl = window.URL.createObjectURL(blob)
                const temLink = document.createElement("a")
                temLink.style.display = "none"
                temLink.href = blobUrl
                const novelTitle = title + ".txt"
                temLink.setAttribute("download", novelTitle)
                document.body.appendChild(temLink)
                temLink.click()
                document.body.removeChild(temLink)
                window.URL.revokeObjectURL(blobUrl)
            })
        },
        async collectionNovel(articleId, type) {


            if (this.$store.state.userInformation.ID === undefined) {
                this.$message.error("请登录!")
            } else {

                const { data: res } = await this.$http(
                    {
                        method: "post",
                        url: "/user/collection/add",
                        headers: { 'Content-Type': 'application/json' },
                        data: {
                            user_id: this.$store.state.userInformation.ID,
                            article_id: new Number(articleId),
                            type: type,
                        }
                    }
                )
                if (res.status === 200) {
                    this.$message.success("收藏成功!")

                } else {
                    this.$message.error(res.message)
                }
            }
        },

    },

    name: "NovelInformation"
}

</script>

<style scoped="scoped">
.novel_infor {
    background-color: rgba(214, 180, 69, 0.493);
    display: flex;
    width: 100%;
    height: 95%;
    flex-wrap: wrap;
}

.el-image {
    width: 200px;
    height: 300px;
}

.infor {
    display: flex;
    width: 1000px;
    height: 260px;
    float: right;
    flex-direction: column;
    align-items: right;
}
</style>