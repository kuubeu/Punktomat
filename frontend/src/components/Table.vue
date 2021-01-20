<template>
  <div class="col-12">
    <b-table
      :fields="fields"
      :items="magazines"
      :select-mode="selectMode"
      ref="selectableTable"
      selectable
      @row-selected="onRowSelected"
    >
      <template #cell(Categories)="data">
        <div class="tags" v-if="data.item.Categories.length < 4">
          <div
            class="chip"
            v-for="category in data.item.Categories"
            v-bind:key="data.item.issn + category"
          >
            {{ category }}
          </div>
        </div>
        <div class="tags" v-else>
          <div
            v-for="(category, index) in data.item.Categories"
            v-bind:key="data.item.issn + category"
          >
            <div class="chip" v-if="index === 0 || index === 1 || index === 2">
              {{ category }}
            </div>
            <div
              class="chip"
              v-if="index === data.item.Categories.length - 1"
              v-bind:id="data.item.issn"
            >
              <b-tooltip
                v-bind:target="data.item.issn"
                triggers="hover"
                placement="bottom"
              >
                {{ data.item.Categories.slice(3).join("\n") }}
              </b-tooltip>
              ...
            </div>
          </div>
        </div>
      </template>
    </b-table>
  </div>
</template>

<script>
export default {
  name: "Table",
  data() {
    return {
      numbers: [...Array(20).keys()],
      selected: [],
    };
  },
  props: ["selectMode", "fields", "magazines"],
  mounted() {
    if (JSON.parse(localStorage.getItem("selected"))) {
      this.selected = [...JSON.parse(localStorage.getItem("selected"))];
      this.selected.map((value) => {
        let i = 0;
        this.magazines.map((valueMagazine) => {
          if (valueMagazine.ID == value.ID) {
            this.$refs.selectableTable.selectRow(i);
          }
          i++;
        });
      });
    }
  },
  methods: {
    onRowSelected(items) {
      this.selected = items;
      localStorage.setItem("selected", JSON.stringify(this.selected));
      this.$emit("clicked");
    },
  },
};
</script>

<style></style>
