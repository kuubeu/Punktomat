import { shallowMount } from "@vue/test-utils";
import App from "../src/App.vue";
import Main from "../src/components/Main.vue";

describe("App", () => {
  const instance = shallowMount(App);
  it("correct rendering of App", () => {
    expect(instance.isVueInstance()).toBeTruthy();
  });
  it("App has components", () => {
    expect(App.components).toBeDefined();
  });
  it("App has once component called main", () => {
    expect(instance.findComponent(Main)).toBeTruthy();
  });
  it("App isn't supposed to have any data", () => {
    expect(App.data).toBeUndefined();
  });
});
