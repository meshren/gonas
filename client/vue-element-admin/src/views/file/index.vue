<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.file_name"
        :placeholder="$t('table.title')"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
      />
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        {{ $t('table.search') }}
      </el-button>
      <el-dropdown class="filter-item" style="margin-left: 10px;" type="primary" @command="handleCommand">
        <el-button type="primary">
          新建
          <i class="el-icon-arrow-down el-icon--right" />
        </el-button>
        <el-dropdown-menu slot="dropdown">
          <el-upload
            class="upload-demo"
            action="/backend/files"
            :on-preview="handlePreview"
            :on-remove="handleRemove"
            multiple
            :data="{directory_id: currentDirectoryID}"
            :on-success="handleSuccess"
            :limit="1"
          >
            <el-dropdown-item command="file">文件</el-dropdown-item>
          </el-upload>
          <el-dropdown-item command="folder">文件夹</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/file/index' }">我的网盘</el-breadcrumb-item>
      <el-breadcrumb-item>活动管理</el-breadcrumb-item>
      <el-breadcrumb-item>活动列表</el-breadcrumb-item>
      <el-breadcrumb-item>活动详情</el-breadcrumb-item>
    </el-breadcrumb>
    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      :cell-style="{cursor: 'pointer'}"
      @sort-change="sortChange"
      @cell-click="handleClick"
    >
      <el-table-column label="名称" prop="display" sortable="custom" align="left" :class-name="getSortClass('id')">
        <template slot-scope="{row}">
          <span>
            <svg-icon :icon-class="handleIcon(row.type)" />
            {{ row.display }}
          </span>
        </template>
      </el-table-column>
      <el-table-column
        label="大小"
        prop="id"
        sortable="custom"
        align="center"
        width="80"
        :class-name="getSortClass('id')"
      >
        <template slot-scope="{row}">
          <span>{{ row.size }}</span>
        </template>
      </el-table-column>
      <el-table-column label="上传时间" width="160px" align="center">
        <template slot-scope="{row}">
          <span>{{ row.created_at | parseTime('{y}-{m}-{d} {h}:{i}:{s}') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.actions')" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="mini" @click="handleUpdate(row)">
            {{ $t('table.edit') }}
          </el-button>
          <el-button v-if="row.deleted_at !== null" size="mini" type="danger" @click="handleDelete(row,$index)">
            {{ $t('table.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"
    />
  </div>
</template>

<script>
import { fetchList, fetchPv, createFolder, updateArticle } from '@/api/file'
import waves from '@/directive/waves' // waves directive
// import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

export default {
  name: 'File',
  components: { Pagination },
  directives: { waves },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    },
    typeFilter(type) {
    }
  },
  data() {
    return {
      currentDirectoryID: 0,
      currentDirectoryName: 0,
      tableKey: 0,
      list: [],
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
        file_name: '',
        directory_id: 0
      },
      folder: {
        display: '',
        parent: ''
      },
      temp: {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        type: '',
        status: 'published'
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    handleSuccess() {
      this.$message('上传成功！')
    },
    handleRemove(file, fileList) {
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log(file)
    },
    handleClick(row, column) {
      if (column.property === 'display' && row.type === 1) {
        this.listQuery.directory_id = row.id
        this.currentDirectoryID = row.id
        this.getList()
      }
    },
    createFolder() {
      this.$prompt('请输入文件夹名', '', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        // inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
        inputErrorMessage: '文件夹名称不合法！'
      }).then(({ value }) => {
        createFolder({ display: value, parent_id: this.currentDirectoryID })
          .then((response) => {
            console.log(response)
            this.$message({
              type: 'success',
              message: '创建文件夹成功: ' + value
            })
          })
          .catch(errors => console.log(errors))
      })
        .catch(errors => console.log(errors))
    },
    handleIcon(type) {
      switch (type) {
        case 0: return 'document'
        case 1: return 'folder'
        case 2: return 'document'
        default: return 'folder'
      }
    },
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.parseData(response.data)
        // this.list = response.data
        // this.total = response.data.length

        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    parseData(data) {
      this.list = []
      const { files, directories } = data
      files.map(file => {
        const { id, display, size, created_at, type } = file
        this.list.push({ id, display, size, created_at, type })
      })
      directories.map(directory => {
        const { id, display, size, created_at } = directory
        this.list.push({ id, display, size, created_at, type: 1 })
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    sortChange(data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID(order) {
      if (order === 'ascending') {
        this.listQuery.sort = '+id'
      } else {
        this.listQuery.sort = '-id'
      }
      this.handleFilter()
    },
    handleCommand(command) {
      console.log(command)
      switch (command) {
        case 'folder':
          console.log('case folder')
          this.createFolder()
          break
        case 'file':
          console.log('case file')
          this.createFile()
      }
    },
    createFile() {},
    handleUpdate(row) {
      this.temp = Object.assign({}, row) // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleDelete(row, index) {
      this.$notify({
        title: '成功',
        message: '删除成功',
        type: 'success',
        duration: 2000
      })
      this.list.splice(index, 1)
    },
    getSortClass: function(key) {
      const sort = this.listQuery.sort
      return sort === `+${key}` ? 'ascending' : 'descending'
    }
  }
}
</script>
