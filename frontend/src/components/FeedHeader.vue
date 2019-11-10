<template>
  <nav class="category-nav nav">
    <header class="list-header">
      <nav role="navigation" class="list-nav">
        <ul class="nav-list left">
          <li class="nav-item active">
            <a href="/welcome/recommended?sort=popular">热门</a>
          </li>
          <li class="nav-item">
            <a href="/welcome/recommended?sort=newest">最新</a>
          </li>
          <li class="nav-item">
            <a href="/welcome/recommended?sort=three_days_hottest">热榜</a>
          </li>
        </ul>
        <div class="dorp-down-area">
          <!---->
        </div>
      </nav>
    </header>
  </nav>
</template>

<script>
export default {};
</script>

<style lang="less" scoped>
.category-nav {
  background-color: #fff;
  display: flex;
  justify-content: space-between;
  margin-top: 25px;
}
.category-nav .list-header {
  width: 100%;
}

.list-header {
  padding: 1.3rem 1rem;
  border-bottom: 1px solid hsla(0, 0%, 59.2%, 0.1);
}

.list-header .left .nav-item {
  padding: 0 1.2rem;
  font-size: 1.16rem;
  border-right: 1px solid hsla(0, 0%, 59.2%, 0.2);
}
.list-header .list-nav {
  display: flex;
  justify-content: space-between;
}
.list-header .nav-list {
  align-items: center;
  line-height: 1;
}

.list-header .list-nav,
.list-header .nav-list {
  display: flex;
  justify-content: space-between;
}
ul {
  padding: 0;
  margin: 0;
}
.category-nav .active {
  color: #007fff;
}
.list-header .left .nav-item {
  padding: 0 1.2rem;
  font-size: 1.16rem;
  border-right: 1px solid hsla(0, 0%, 59.2%, 0.2);
}
.nav-item {
  position: relative;
  cursor: pointer;
}

li {
  list-style: none;
}
.nav-item > a:before {
  content: "";
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
}
.list-header .nav-list .nav-item.active a,
.list-header .nav-list .nav-item a:hover {
  color: #007fff;
}
a {
  text-decoration: none;
  cursor: pointer;
  color: #909090;
  margin: initial;
  font-size: 15px;
}
</style>