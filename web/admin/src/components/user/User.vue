<template>
  <div class="users">

    <el-dialog title="编辑用户" v-model="centerDialogVisible" width="30%" destroy-on-close center>
      <el-form :inline="true" :model="editUser" label-width="100px" class="demo-form-inline">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="editUser.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <span>是否设为管理员</span>
          <el-select v-model="editUser.role" size="mini" :label="editUser.role">
            <el-option value="是"><span>是</span>
            </el-option>
            <el-option value="否"><span>否</span>
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="centerDialogVisible = cancelBut(editUser)">取 消</el-button>
          <el-button type="primary" @click="editSubmit">确 定</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog title="新增用户" v-model="addDialogVisible" width="30%" destroy-on-close center>
      <el-form :inline="true" :rules="rules" ref="addUser" :model="addUser" label-width="100px"
        class="demo-form-inline">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="addUser.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="addUser.password" autocomplete="off" show-password></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="checkPassWord">
          <el-input v-model="addUser.checkPassWord" autocomplete="off" show-password></el-input>
        </el-form-item>
        <el-form-item>
          <span>是否设为管理员</span>
          <el-select v-model="addUser.role" size="mini" label="addUser.role">
            <el-option value="是"><span>是</span>
            </el-option>
            <el-option value="否"><span>否</span>
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="addDialogVisible = cancelBut(addUser)">取 消</el-button>
          <el-button type="primary" @click="addUserSubmit">确 定</el-button>
        </span>
      </template>
    </el-dialog>

    <el-table :data="tableData.slice((currentPage - 1) * PageSize, currentPage * PageSize)" class="tableContent"
      style="min-height:750px;" @selection-change="handleSelectionChange" align="center">
      <el-table-column type="index" :index="indexMethod"> </el-table-column>
      <el-table-column label="创建时间" prop="CreatedAt"> </el-table-column>
      <el-table-column label="用户名" prop="username"> </el-table-column>
      <el-table-column label="密码" prop="password"> </el-table-column>
      <el-table-column label="类型" prop="role"> </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button icon="delete" size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
          <el-button icon="edit" size="mini" type="primary" @click="editUserFun(scope.row)">编辑</el-button>
        </template>
      </el-table-column>
      <el-table-column align="right">
        <template #header>
          <el-input v-model="editUsername" clearable size="mini" placeholder="输入用户名搜索" @clear="editUserClear"
            @keyup.enter.native="searchUser" />
        </template>
      </el-table-column>
    </el-table>

    <el-button icon="Plus" type="danger" style="float: left;" @click="addUserBut">添加用户</el-button>

    <el-container class="page" style="float: right;">
      <el-button type="primary" @click="jumpFirstPage" class="my-btn" style="float: left;">首页</el-button>
      <el-pagination ref="pagination" background prev-text="上一页" next-text="下一页" layout="total,prev, pager, next, slot"
        @current-change="handleCurrentChange" :current-page="currentPage" :total="total" :page-size="PageSize">
      </el-pagination>
      <el-button type="primary" @click="jumpLastPage" class="my-btn" style="float: right;">尾页</el-button>
    </el-container>
  </div>
</template>

<script>

export default {
  data() {
    var ISPWD = /^(?=.*[0-9])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*,\._\+(){}])[0-9a-zA-Z!@#$%^&*,\\._\+(){}]{8,20}$/;
    // 密码校验
    const validatePassword = (rule, value, callback) => {
      if (!ISPWD.test(this.addUser.password)) {
        callback(new Error("用户密码必须包含大写字母、小写字母、数字和特殊符号"));
      } else {
        callback();
      }
    }

    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.addUser.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      tableData: [],
      search: '',
      multipleSelection: [],
      editUsername: '',
      total: 0,
      currentPage: 1,
      PageSize: 10,
      centerDialogVisible: false,
      addDialogVisible: false,
      addUser: {
        username: '',
        password: '',
        checkPassWord: '',
        role: '',
      },
      editUser: {
        id: 0,
        username: '',
        role: '',
      },
      rules: {
        username: [
          { required: true, trigger: "blur", message: "请输入您的账号" },
          { min: 2, max: 20, message: '用户账号长度必须介于 2 和 20 之间', trigger: 'blur' }
        ],
        password: [
          { required: true, trigger: "blur", message: "请输入您的密码" },
          { min: 8, max: 20, message: '用户密码长度必须介于 8 和 20 之间', trigger: 'blur' },
          { required: true, validator: validatePassword, trigger: 'blur' }
        ],
        checkPassWord: [{ validator: validatePass2, trigger: 'blur' }],
      },
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('/users?role=1')
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.tableData = res.data
      this.total = res.total
    },
 
    async handleDelete(row) {
      const confirmR = await this.$confirm(
        "此操作将删除该用户, 是否继续?",
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
        const { data: res } = await this.$http({
          method: "delete",
          url: "/user/delete/" + row.ID,
          headers: { 'Content-Type': 'multipart/form-data' },
        })

        if (res.status === 200) {
          this.$message.success("删除成功")
          this.getUserList()
        } else if (res.status === 1004 || 1005 || 1006 || 1007) {
          this.$message.error("非法操作")
          this.$router.push('/login')
        } else {
          this.$message.error(res.message)
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

    editUserFun(row) {
      this.centerDialogVisible = true
      this.editUser.id = row.ID
    },

    async editSubmit() {
      if (this.editUser.role === "是") {
        const { data: res } = await this.$http({
          method: "put",
          url: "/user/update/" + this.editUser.id,
          headers: { 'Content-Type': 'application/json' },
          data: {
            username: this.editUser.username,
            role: 1
          }
        })
        if (res.status === 200) {
          this.$message.success("修改成功")
          this.editUser = {
            id: 0,
            username: '',
            role: '',
          }
          this.centerDialogVisible = false
          this.getUserList()
        } else if (res.status === 1001) {
          this.$message.error("用户名已存在")
        }

      } else if (this.editUser.role === "") {
        this.$message.error("未设置用户类型")
      } else {
        const { data: res } = await this.$http({
          method: "put",
          url: "/user/update/" + this.editUser.id,
          headers: { 'Content-Type': 'application/json' },
          data: {
            username: this.editUser.username,
            role: 0
          }
        })
        if (res.status === 200) {
          this.$message.success("修改成功")
          this.$refs[editUser].resetFields();
          this.editUser = {
            id: 0,
            username: '',
            role: '',
          }
          this.centerDialogVisible = false
          this.getUserList()
        } else if (res.status === 1001) {
          this.$message.error("用户名已存在")
        }
      }
    },

    async searchUser() {
      if (this.editUsername === "") {
        this.getUserList()
      } else {
        const { data: res } = await this.$http({
          method: "get",
          url: "/user?username=" + this.editUsername,
        })
        if (res.status === 200) {
          this.tableData = [res.data]
          this.total = 1
        } else {
          this.$message.error(res.message)
        }
      }
    },
    editUserClear() {
      this.editUsername = ''
    },
    addUserBut() {
      this.addDialogVisible = true
    },
    addUserSubmit() {
      this.$refs.addUser.validate(async (valid) => {
        if (!valid || this.addUser.role === "") return this.$message.error('输入非法数据，请重新输入')
        if (this.addUser.role === "是") {
          this.addUser.role = 1
        } else {
          this.addUser.role = 0
        }
        const { data: res } = await this.$http({
          method: "post",
          url: "/user/add",
          headers: {
            'Content-Type': 'text/plain',
          },
          data: {
            username: this.addUser.username,
            password: this.addUser.password,
            role: this.addUser.role,
          },
        })

        if (res.status === 200) {
          this.getUserList()
          Object.keys(this.addUser).forEach((key) => (this.addUser[key] = ''))
          this.addDialogVisible = false
          this.$message.success("添加成功")
        } else {
          this.$message.error(res.message)
        }
      })
    },
    cancelBut(userTable) {
      Object.keys(userTable).forEach((key) => (userTable[key] = ''))
      return false
    },
  },
  name: "User",
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