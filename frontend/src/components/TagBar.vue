<template>
  <div class="tag-bar">
    <ul class="nav-list left">
      <li class="nav-item active">
        <div class="category-popover-box">
          <!---->
          <a>推荐</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>后端</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>前端</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>Android</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>iOS</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>人工智能</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>开发工具</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>代码人生</a>
        </div>
      </li>
      <li class="nav-item">
        <div class="category-popover-box">
          <!---->
          <a>阅读</a>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
export default {};
</script>

<style lang="less" scoped>
.tag-bar{
    position: fixed;
    top: 60px;
    width: 100%;
    height: 3.833rem;
    z-index: 100;
    background-color:  rgb(255, 255, 255);
}
.nav-list {
    max-width: 960px;
    height: 100%;
    margin: auto;
    display: flex;
    align-items: center;
    line-height: 1;
}
.nav-list .nav-item:first-child {
    padding: 0 9px 0 0;
}
.nav-list .nav-item {
    height: 100%;
    align-items: center;
    display: flex;
    flex-shrink: 0;
    color: #71777c;
    padding: 0 9px;
    font-size: 15px;
}
.nav-item {
    position: relative;
    cursor: pointer;
}

li {
    list-style: none;
}
:hover{
    color: #007fff;
}
</style>