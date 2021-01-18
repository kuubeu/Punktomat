<template>
  <div class=".container-fluid">
    <b-navbar variant="faded" type="light">
      <b-navbar-brand tag="h1" class="mb-0">Punktomat</b-navbar-brand>
    </b-navbar>
    <div class="row justify-content-md-center">
      <Table
        v-bind:selectMode="selectMode"
        v-bind:loading="loading"
        v-bind:fields="fields"
        v-bind:magazines="magazines"
      />
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Table from "./Table";

export default {
  name: "Main",
  components: { Table },
  data() {
    return {
      magazines: null,
      loading: true,
      fields: [
        "ID",
        "title",
        "issn",
        "points",
        { key: "Categories", label: "Kategorie" },
      ],
      selectMode: "multi",
    };
  },
  mounted() {
    axios.get("http://127.0.0.1:4000/api/scienceMagazine").then((response) => {
      this.magazines = response.data.results;
      this.total = response.data.total;
      this.loading = false;
    });
  },
};
</script>

<style>
.chip {
  -webkit-align-items: center;
  align-items: center;
  -webkit-border-radius: 15px;
  border-radius: 15px;
  -webkit-box-shadow: inset 0 0 0 1px rgb(100 121 143 / 30%);
  box-shadow: inset 0 0 0 1px rgb(100 121 143 / 30%);
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  color: #5f6368;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  height: 30px;
  margin-bottom: 4px;
  margin-right: 8px;
  /* max-width: 160px; */
  overflow: hidden;
  padding: 0 12px;
  white-space: nowrap;
}
.tags {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  min-height: 30px;
  padding-bottom: 2px;
  margin-top: 8px;
  max-width: calc(100vw - 184px);
  width: 100%;
  flex-wrap: wrap;
}
</style>
