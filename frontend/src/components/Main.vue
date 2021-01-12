<template>
  <div class=".container-fluid">
    <b-navbar variant="faded" type="light">
      <b-navbar-brand tag="h1" class="mb-0">Punktomat</b-navbar-brand>
    </b-navbar>
    <div class="row justify-content-md-center">
      <div class="col-11">
        <b-table
          :fields="fields"
          :items="magazines"
          :select-mode="selectMode"
          ref="selectableTable"
          selectable
          @row-selected="onRowSelected"
        >
          <template #cell(ID)="data">
            {{ data.item.ID }}
          </template>
          <template #cell(points)="data">
            {{ data.item.points }}
          </template>
          <template #cell(title)="data">
            {{ data.item.title }}
          </template>
          <template #cell(issn)="data">
            {{ data.item.issn }}
          </template>
          <template #cell(Categories)="data">
            <div class="tags">
              <div
                class="chip"
                v-for="category in data.item.Categories"
                v-bind:key="data.item.issn + category"
              >
                {{ category }}
              </div>
            </div>
          </template>
        </b-table>
      </div>
    </div>
    {{ selected }}
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Main",
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
      numbers: [...Array(20).keys()],
      selected: [],
    };
  },
  methods: {
    onRowSelected(items) {
      this.selected = items;
    },
  },
  mounted() {
    setTimeout(
      () =>
        axios.get("http://127.0.0.1:4000").then((response) => {
          this.magazines = response.data;
          this.loading = false;
        }),
      2000
    );
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
