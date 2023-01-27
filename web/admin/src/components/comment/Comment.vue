<template>
    <div class="users">
      <el-table :data="tableData.filter(data => !search || data.address.toLowerCase().includes(search.toLowerCase()))"
        class="tableContent" style="min-height:750px;" @selection-change="handleSelectionChange">
        <el-table-column type="index" :index="indexMethod"> </el-table-column>
        <el-table-column label="文章名" prop="name"> </el-table-column>
        <el-table-column label="类型" prop="address"> </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="mini" type="danger" @click="handleDelete(scope.$index, tableData)">删除</el-button>
            <el-button size="mini" type="primary" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
          </template>
        </el-table-column>
        <el-table-column align="right">
          <template #header>
            <el-input v-model="search" size="mini" placeholder="输入关键字搜索" />
          </template>
        </el-table-column>
      </el-table>
  
      <el-container class="page" style="float: right;">
        <el-button type="primary" @click="jumpFirstPage" class="my-btn" style="float: left;">首页</el-button>
        <el-pagination ref="pagination" background prev-text="上一页" next-text="下一页" layout="prev, pager, next, slot"
          @current-change="handleCurrentChange" :current-page="currentPage" :total="1000" :page-size="5">
        </el-pagination>
        <el-button type="primary" @click="jumpLastPage" class="my-btn" style="float: right;">尾页</el-button>
      </el-container>
    </div>
  </template>
    
  <script>
  export default {
    data() {
      return {
        tableData: [],
        tableAllData: [
        ],
  
        search: '',
        multipleSelection: [],
      }
    },
    methods: {
  
      handleEdit(index, row) {
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
          // 进行删除
          row.splice(index, 1)
          his.$message.success("删除成功");
  
        } else {
          return this.$message.info("当前操作已取消");
        }
      },
  
      indexMethod(index) {
        return index + 1;
      },
  
      setCurrent(rows) {
        rows.forEach((row) => {
          console.log(row)
        })
      },
  
      handleSelectionChange(val) {
        this.multipleSelection = val;
      },
      //todo 索引
      handleDeleteAll() {
        let _this = this;
        if (this.multipleSelection.length == 0) {
          this.$message.error("您还未选中用户！");
        } else {
          this.$confirm("此操作将永久删除用户, 是否继续?", "提示", {
            confirmButtonText: "确定",
            cancelButtonText: "取消",
            type: "warning"
          }).then(() => {
            _this.multipleSelection.forEach(item => {
              console.log(item)
              this.tableData.splice(item, 1);
            });
  
            this.$message({
              type: "success",
              message: "删除成功!"
            });
          });
        }
      },
  
      jumpFirstPage() {
        this.tableAllData.forEach(data => {
          if (this.tableData.length < 5) {
            this.tableData.push(data)
          }
        });
      },
      handleCurrentChange() {
        this.tableData.forEach((data, index) => {
          this.tableData[index] = this.tableData[index + 1]
        })
      },
  
    },
    name: "Comment",
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