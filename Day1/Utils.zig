const std = @import("std");
const Allocator = std.mem.Allocator;
const ArrayList = std.ArrayList;

pub fn readFileIntoArr(allocator: Allocator, filepath: []const u8) ![]u8 {
    const file = try std.fs.cwd().openFile(
        filepath,
        .{ .mode = .read_only },
    );
    defer file.close();

    const fileStats = try file.stat();
    const fileSize = fileStats.size;

    var reader = file.reader();

    const data: []u8 = try reader.readAllAlloc(allocator, fileSize * 2);
    
    return data;
}


pub fn abs(comptime T: type, num: T) T {
    if (num < 0) {
        return num * -1;
    }
    else return num;
}


pub fn printArr(comptime T: type, arr: []T) !void {
    const writer = std.io.getStdOut().writer();

    for (arr) |x| {
        try writer.print("{d} \n", .{x});
    }
}


pub fn printAList(comptime T: type, list: ArrayList(T)) !void {
    const writer = std.io.getStdOut().writer();

    for (list.items) |x| {
        try writer.print("{d} \n", .{x});
    }
}