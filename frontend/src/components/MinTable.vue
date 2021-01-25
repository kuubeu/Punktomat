<template>
  <v-data-table
    :items="selectedMagazines"
    :headers="headers"
    class="elevation-1"
  >
    <template v-slot:[`item.points`]="{ item }">
      <v-layout justify-center>
        <v-chip :color="getColor(item.points)" dark>
          {{ item.points }}
        </v-chip>
      </v-layout>
    </template>
    <template v-slot:[`item.delete`]="{ item }">
      <v-layout justify-center>
        <v-icon small @click="deleteItem(item)">
          mdi-delete
        </v-icon>
      </v-layout>
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
      console.log(item);
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
          width: "clamp(100px, 30vw, 250px)",
        },
        {
          text: "Punkty",
          value: "points",
          sortable: false,
          width: "88px",
        },
        {
          text: "ISSN",
          value: "issn",
          sortable: false,
          width: "100px",
        },
        { text: "Usuń", value: "delete", sortable: false, width: "20px" },
      ],
    };
  },
};
</script>

<style></style>
