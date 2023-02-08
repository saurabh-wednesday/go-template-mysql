// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrations)
	t.Run("Posts", testPosts)
	t.Run("Roles", testRoles)
	t.Run("RolesTests", testRolesTests)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsDelete)
	t.Run("Posts", testPostsDelete)
	t.Run("Roles", testRolesDelete)
	t.Run("RolesTests", testRolesTestsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsQueryDeleteAll)
	t.Run("Posts", testPostsQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("RolesTests", testRolesTestsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceDeleteAll)
	t.Run("Posts", testPostsSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("RolesTests", testRolesTestsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsExists)
	t.Run("Posts", testPostsExists)
	t.Run("Roles", testRolesExists)
	t.Run("RolesTests", testRolesTestsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsFind)
	t.Run("Posts", testPostsFind)
	t.Run("Roles", testRolesFind)
	t.Run("RolesTests", testRolesTestsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsBind)
	t.Run("Posts", testPostsBind)
	t.Run("Roles", testRolesBind)
	t.Run("RolesTests", testRolesTestsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsOne)
	t.Run("Posts", testPostsOne)
	t.Run("Roles", testRolesOne)
	t.Run("RolesTests", testRolesTestsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsAll)
	t.Run("Posts", testPostsAll)
	t.Run("Roles", testRolesAll)
	t.Run("RolesTests", testRolesTestsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsCount)
	t.Run("Posts", testPostsCount)
	t.Run("Roles", testRolesCount)
	t.Run("RolesTests", testRolesTestsCount)
	t.Run("Users", testUsersCount)
}

func TestInsert(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsInsert)
	t.Run("GorpMigrations", testGorpMigrationsInsertWhitelist)
	t.Run("Posts", testPostsInsert)
	t.Run("Posts", testPostsInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("RolesTests", testRolesTestsInsert)
	t.Run("RolesTests", testRolesTestsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("PostToUserUsingAuthor", testPostToOneUserUsingAuthor)
	t.Run("UserToRoleUsingRole", testUserToOneRoleUsingRole)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("RoleToUsers", testRoleToManyUsers)
	t.Run("UserToAuthorPosts", testUserToManyAuthorPosts)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("PostToUserUsingAuthorPosts", testPostToOneSetOpUserUsingAuthor)
	t.Run("UserToRoleUsingUsers", testUserToOneSetOpRoleUsingRole)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("PostToUserUsingAuthorPosts", testPostToOneRemoveOpUserUsingAuthor)
	t.Run("UserToRoleUsingUsers", testUserToOneRemoveOpRoleUsingRole)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("RoleToUsers", testRoleToManyAddOpUsers)
	t.Run("UserToAuthorPosts", testUserToManyAddOpAuthorPosts)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("RoleToUsers", testRoleToManySetOpUsers)
	t.Run("UserToAuthorPosts", testUserToManySetOpAuthorPosts)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("RoleToUsers", testRoleToManyRemoveOpUsers)
	t.Run("UserToAuthorPosts", testUserToManyRemoveOpAuthorPosts)
}

func TestReload(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsReload)
	t.Run("Posts", testPostsReload)
	t.Run("Roles", testRolesReload)
	t.Run("RolesTests", testRolesTestsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsReloadAll)
	t.Run("Posts", testPostsReloadAll)
	t.Run("Roles", testRolesReloadAll)
	t.Run("RolesTests", testRolesTestsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSelect)
	t.Run("Posts", testPostsSelect)
	t.Run("Roles", testRolesSelect)
	t.Run("RolesTests", testRolesTestsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsUpdate)
	t.Run("Posts", testPostsUpdate)
	t.Run("Roles", testRolesUpdate)
	t.Run("RolesTests", testRolesTestsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceUpdateAll)
	t.Run("Posts", testPostsSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("RolesTests", testRolesTestsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
