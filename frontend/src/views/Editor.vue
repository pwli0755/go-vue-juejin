<template>
  <div id="editor">
    <mavon-editor @save='save' v-model="value" style="height: 100%"/>
</div>
</template>

<script>
export default {
    data(){
        return {
            value:'',
        }
    },
    methods:{
        save(){
            // 保存至
            localStorage.data = JSON.stringify(arguments);
            console.log(localStorage.getItem('data'))
        },
    }
}
</script>

<style lang='less' scoped>
#editor {
        margin: auto;
        height: 800px;
    }

</style>