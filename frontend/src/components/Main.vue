<template>
  <div class=".container-fluid">
    <b-navbar variant="faded" type="light">
      <b-navbar-brand tag="h1" class="mb-0">Punktomat</b-navbar-brand>
      <md-badge
        class="md-primary"
        v-bind:md-content="selected.length"
        md-dense
        v-if="selected"
      >
        <md-button @click="showModal" class="md-icon-button">
          <md-icon>article</md-icon>
        </md-button>
      </md-badge>
      <md-icon v-else>article</md-icon>
    </b-navbar>
    <div class="row justify-content-md-center">
      <b-skeleton-table
        v-if="loading"
        :rows="10"
        :columns="5"
        :table-props="{ striped: true }"
      ></b-skeleton-table>
      <Table
        v-else
        @clicked="getSelected()"
        v-bind:selectMode="selectMode"
        v-bind:loading="loading"
        v-bind:fields="fields"
        v-bind:magazines="magazines"
        v-bind:selected="selected"
      />
    </div>
    <b-modal
      ref="my-modal"
      size="xl"
      hide-footer
      title="Using Component Methods"
    >
      <div
        class="d-block text-center"
        v-for="select in selected"
        v-bind:key="select.issn"
      >
        <p>{{ select }}</p>
      </div>
      <b-button class="mt-3" variant="outline-danger" block @click="hideModal"
        >Zamknij</b-button
      >
      <b-button class="mt-2" variant="outline-warning" block @
        >Generuj PDF</b-button
      >
    </b-modal>
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
      selected: [],
    };
  },
  methods: {
    getSelected: function() {
      this.selected = [...JSON.parse(localStorage.getItem("selected"))];
    },
    showModal() {
      this.$refs["my-modal"].show();
    },
    hideModal() {
      this.$refs["my-modal"].hide();
    },
  },
  mounted() {
    if (JSON.parse(localStorage.getItem("selected"))) {
      this.selected = [...JSON.parse(localStorage.getItem("selected"))];
    }
    axios.get(`${process.env.VUE_APP_API_URL}/scienceMagazine`).then((response) => {
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
