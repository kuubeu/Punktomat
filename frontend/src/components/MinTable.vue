<template>
  <v-data-table
    :items="selectedMagazines"
    :headers="headers"
    class="elevation-0"
  >
    <template v-slot:[`item.title`]="{ item }">
      <v-layout class="mt-2 mb-2">
        {{ item.title }}
      </v-layout>
    </template>
    <template v-slot:[`item.points`]="{ item }">
      <v-chip :color="getColor(item.points)" dark>
        {{ item.points }}
      </v-chip>
    </template>
    <template v-slot:[`item.delete`]="{ item }">
      <v-layout justify-center>
        <v-icon @click="deleteItem(item)">
          mdi-delete
        </v-icon>
      </v-layout>
    </template>
    <template v-slot:[`header.delete`]>
      <v-tooltip
        left
        :open-delay="200"
        transition="fade-transition"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            v-bind="attrs"
            v-on="on"
            color="error"
            @click="deleteAllItems"
            depressed
            icon
          >
            <v-icon>
              mdi-delete-alert
            </v-icon>
          </v-btn>
        </template>
        <span>Wyczyść listę</span>
      </v-tooltip>
    </template>
  </v-data-table>
</template>

<script>
export default {
  props: ["selectedMagazines"],
  methods: {
    getColor(item) {
      if (item > 100) return "green";
      else if (item > 50) return "orange";
      else return "red";
    },
    deleteItem(item) {
      const index = this.selectedMagazines.indexOf(item);
      this.selectedMagazines.splice(index, 1);
    },
    deleteAllItems() {
      this.selectedMagazines.splice(0, this.selectedMagazines.length);
      console.log("zdarzenie");
    },
  },
  data() {
    return {
      editedIndex: -1,
      editedItem: {},
      headers: [
        {
          text: "Tytuł",
          value: "title",
          sortable: false,
        },
        {
          text: "Punkty",
          value: "points",
          sortable: false,
          width: "80px",
          align: "center",
        },
        {
          text: "ISSN",
          value: "issn",
          sortable: false,
          width: "100px",
        },
        {
          text: "Usuń",
          value: "delete",
          sortable: false,
          width: "20px",
          align: "right",
        },
      ],
    };
  },
};
</script>

<style></style>
