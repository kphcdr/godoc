<template>
  <div class="hello">
    <Header> </Header>

    <el-container>
          <el-card class="center-card">
            <el-form  status-icon  label-width="0px" class="demo-ruleForm">
              <h2></h2>
              <el-form-item label="" >
              <el-radio v-model="item_type" label="1">{{$t('item_type1')}}</el-radio>

              </el-form-item>


              <el-form-item label="" >
                <el-input type="text" auto-complete="off" :placeholder="$t('item_name')" v-model="item_name"></el-input>
              </el-form-item>

              <el-form-item label="" >
                <el-input type="text" auto-complete="off" :placeholder="$t('item_description')" v-model="item_description"></el-input>
              </el-form-item>

               <el-form-item label="" >
                <el-button type="primary" style="width:100%;" @click="onSubmit" >{{$t('submit')}}</el-button>
              </el-form-item>

              <el-form-item label=""  >
                  <router-link to="/item/index">{{$t('goback')}}</router-link>
              </el-form-item>
            </el-form>
          </el-card>
    </el-container>

    <Footer> </Footer>

  </div>
</template>

<script>


export default {
  name: 'Login',
  components : {

  },
  data () {
    return {
      item_type: '1',
      item_name: '',
      item_description:'',
      item_domain: '',
      show_copy:false,
      itemList:{},
      copy_item_id:""

    }

  },
  methods: {
    get_item_list(){
        var that = this ;
        var url = DocConfig.server+'/api/item/myList';

        var params = new URLSearchParams();

        that.axios.get(url, params)
          .then(function (response) {
            if (response.data.error_code === 0 ) {
              //that.$message.success("加载成功");
              var json = response.data.data ;
              that.itemList = json ;
            }else{
              that.$alert(response.data.error_message);
            }

          });
    },
    choose_copy_item(item_id){
      for (var i = 0; i < this.itemList.length; i++) {
        if (item_id == this.itemList[i].item_id) {
            this.item_name = this.itemList[i].item_name+'--copy';
            this.item_description = this.itemList[i].item_description;
        };
      };
    },
    onSubmit() {
        var that = this ;
        var url = DocConfig.server+'/api/item/add';

        var params = new URLSearchParams();
        params.append('item_type', this.item_type);
        params.append('item_name', this.item_name);
        params.append('password', this.password);
        params.append('item_domain', this.item_domain);
        params.append('copy_item_id', this.copy_item_id);
        params.append('item_description', this.item_description);

        that.axios.post(url, params)
          .then(function (response) {
            if (response.data.error_code === 0 ) {
              that.$router.push({path:'/item/index'});
            }else{
              that.$alert(response.data.error_message);
            }

          });
      },
  },
  mounted() {
    this.get_item_list();
    /*给body添加类，设置背景色*/
    document.getElementsByTagName("body")[0].className="grey-bg";
  },
  beforeDestroy(){
    /*去掉添加的背景色*/
    document.body.removeAttribute("class","grey-bg");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.center-card a {
  font-size: 12px;
}

.center-card{
  text-align: center;
  width: 350px;
}

</style>
