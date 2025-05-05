
db = db.getSiblingDB("golang_test");

const time = Date.now()

// 1. Insert user
const user = {
  _id:"usr_fb4a15c6-3fe9-488a-b853-bc869d37eb45",
  username: "admin",
  email: "admin@admin.com",
  phone: "",
  fullname: "super admin",
  birthday: "",
  password: "$2a$10$Pnmvw2eudCsIhlUtL6ojwuoJbPciniUNMbwHnhHCTMzbOxDtIyji2",
  ctime: time,
  mtime: time,
  dtime: 0
};
db.users.insertOne(user);

const groupRole = {
  _id:"gr_67a17d82-1b0a-4412-b4f3-0b0709de586c",
  name:  "admin",
  description: "super admin",
  status:  "status_type_enabled",
  member_ids: ["usr_fb4a15c6-3fe9-488a-b853-bc869d37eb45"],
  ctime: time,
  mtime: time,
  dtime: 0
};
db.group_roles.insertOne(groupRole);

const feature = {
  _id: "ft_435816b9-be0d-4ece-8ea6-9a9917548650",
  name: "Feature",
  key_menu:  "/features",
  url_view:  "/features",
  api:  "/features",
  action: ["/features [post]", "/features [get]", "/features/:id [get]", "/features/:id [put]", "/features/:id [delete]"],
  status: "status_type_enabled",
  role_names: ["admin"],
  ctime:time,
  mtime: time,
  dtime: 0
};
db.features.insertOne(feature);

const roleF = {
  _id: "ft_05982bf7-cd23-4723-a1ed-8e291204c272",
  name: "GroupRole",
  key_menu:  "/group-roles",
  url_view:  "/group-roles",
  api:  "/group-roles",
  action: ["/group-roles [post]", "/group-roles [get]", "/group-roles/:id [get]", "/group-roles/:id [put]", "/group-roles/:id [delete]"],
  status: "status_type_enabled",
  role_names: ["admin"],
  ctime:time,
  mtime: time,
  dtime: 0
};
db.features.insertOne(roleF);

const uF = {
  _id: "ft_7a6d63dd-23f8-4f29-8302-0590a9bc3864",
  name: "User",
  key_menu:   "/users",
  url_view:   "/users",
  api:   "/users",
  action: ["/users [post]", "/users [get]", "/users/:id [get]", "/users/:id [delete]"],
  status: "status_type_enabled",
  role_names: ["admin"],
  ctime:time,
  mtime: time,
  dtime: 0
};
db.features.insertOne(uF);
