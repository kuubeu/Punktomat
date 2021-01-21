import { createLocalVue, mount, shallowMount, render } from "@vue/test-utils";
import Table from "../src/components/Table.vue";

describe("Table", () => {
  const instance = shallowMount(Table, {
    propsData: {
      loading: true,
      magazines: [],
      totalMagazines: 10,
    },
  });
  it("crorrect rendering table", () => {
    expect(instance.exists()).toBe(true);
  });
  it("correct rendering of Table vue", () => {
    expect(instance.isVueInstance()).toBeTruthy();
  });

  it("Table hasn't components", () => {
    expect(Table.components).toBeUndefined();
  });
  it("Table is supposed to have any data", () => {
    expect(Table.data).toBeDefined();
  });
  it("Table props to be defined", () => {
    expect(Table.props).toBeDefined();
  });
  it("Skieleton test", () => {
    expect(instance.find(".v-skeleton-loader")).toBeTruthy();
  });
});
