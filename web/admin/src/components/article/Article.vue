<template>
  <div class="users">
    <el-col>
      <el-input placeholder="输入小说名搜索" style="width: 15%;float: left;" v-model.trim="artTitle"
        v-on:keyup.enter="searchArt"></el-input>
      <el-button type="primary" style="float: left;" @click="searchArt">搜索</el-button>
      <label style="float: right;">选择类型</label>
      <el-select size="mini" style="float: right;width: 10%;" v-model="optType" label="optType">
        <el-option value="玄幻" @click.native="getArticleList(optType)"><span>玄幻</span></el-option>
        <el-option value="武侠" @click.native="getArticleList(optType)"><span>武侠</span></el-option>
        <el-option value="网游" @click.native="getArticleList(optType)"><span>网游</span></el-option>
        <el-option value="穿越" @click.native="getArticleList(optType)"><span>穿越</span></el-option>
        <el-option value="科幻" @click.native="getArticleList(optType)"><span>科幻</span></el-option>
        <el-option value="悬疑" @click.native="getArticleList(optType)"><span>悬疑</span></el-option>
        <el-option value="其他" @click.native="getArticleList(optType)"><span>其他</span></el-option>
      </el-select>
    </el-col>

    <el-table :data="tableData.slice((currentPage - 1) * PageSize, currentPage * PageSize)" class="tableContent"
      style="min-height:750px;" @selection-change="handleSelectionChange">
      <el-table-column type="index" :index="indexMethod"> </el-table-column>
      <el-table-column label="类型" prop="type"> </el-table-column>
      <el-table-column label="文章名" prop="title"> </el-table-column>
      <el-table-column label="作者" prop="author"> </el-table-column>
      <el-table-column label="图片">
        <template v-slot="scope">
          <el-image :src="scope.row.img" :preview-src-list="[scope.row.img]" :preview-teleported="true" :fit="fit"
            style="width: 50px;height:50px;cursor:pointer" fit="fill">
          </el-image>
        </template>
      </el-table-column>

      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-container class="page" style="float: right;">
      <el-button type="primary" @click="jumpFirstPage" style="float: left;">首页</el-button>
      <el-pagination ref="pagination" background prev-text="上一页" next-text="下一页" layout="total,prev, pager, next,jumper"
        @size-change="handleSizeChange" @current-change="handleCurrentChange" :total="total" :current-page="currentPage"
        :page-size="PageSize">
      </el-pagination>
      <el-button type="primary" @click="jumpLastPage" style="float: right;">尾页</el-button>
    </el-container>
  </div>
</template>
  
<script>
export default {
  data() {
    return {
      jumpPage: 1,
      artTitle: '',
      tableData: [],
      optType: '玄幻',
      total: 0,
      currentPage: 1,
      PageSize: 10,

    }
  },
  created() {
    this.getArticleList(this.optType)
  },
  methods: {
    async getArticleList(articleType) {
      const url = '/article/type?type=' + articleType
      const { data: res } = await this.$http.get(url)
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }

      this.tableData = res.articles
      this.total = res.total
    },

    async handleDelete(index, row) {
      const confirmR = await this.$confirm(
        "此操作将删除该数据, 是否继续?",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
          center: true,
        }
      ).catch((err) => err);
      if (confirmR == "confirm") {
        
      
        const { data: res } = await this.$http({
          method: "delete",
          url: "/article/delete",
          headers: { 'Content-Type': 'application/json' },
          data: {
            "id": row.id,
            "type": row.type
          }
        })
        if (res.status === 200) {
          this.$message.success("删除成功！")
        } else {
          this.$message.error("删除失败！")
        }

      } else {
        return this.$message.info("当前操作已取消");
      }
    },

    indexMethod(index) {
      return index + 1;
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
    async searchArt() {
      if (this.artTitle !== "") {
        const { data: res } = await this.$http(
          {
            method: "post",
            url: "/article",
            headers: { 'Content-Type': 'application/json' },
            data: {
              title: this.artTitle,
            }
          }
        )
        if (res.status !== 200) {
          this.$message.error(res.message)
          this.tableData = []
          this.total = 0
        } else {
          this.total = 1
          this.tableData = [res.article]
        }

      } else {
        this.getArticleList(this.optType)
      }
    },
    JhandleSizeChange(val) {
      this.currentPage = val
    },
  },
  name: "Article",
}
</script>
  
<style scoped>
.users {
  width: 100%;
  min-width: 1200px;
  height: 750px;
}

.tableContent {
  width: 100%;
  line-height: 50px;
}
</style>