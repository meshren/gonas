<template>
  <el-upload :data="dataObj" :multiple="true" :on-success="handleSuccess" action="/backend/files" drag>
    <i class="el-icon-upload" />
    <div class="el-upload__text">
      将文件拖到此处，或<em>点击上传</em>
    </div>
  </el-upload>
</template>

<script>
// import { getToken } from '@/api/qiniu'
// 获取七牛token 后端通过Access Key,Secret Key,bucket等生成token
// 七牛官方sdk https://developer.qiniu.com/sdk#official-sdk

export default {
  data() {
    return {
      dataObj: { token: '', key: '' },
      image_uri: [],
      fileList: []
    }
  },
  created() {

  },
  methods: {
    handleSuccess(response) {
      const { message } = response
      this.$message(message)
    }
  }
}
</script>
