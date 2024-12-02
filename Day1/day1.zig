const std = @import("std");
const Allocator = std.mem.Allocator;
const expect = std.testing.expect;
const ArrayList = std.ArrayList;

const sort = std.mem.sort;
const dbg = std.debug;

const LType: type = i64;

pub fn main() !void {
    const writer = std.io.getStdOut().writer();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    //? Get file
    const file = try std.fs.cwd().openFile(
        "Data.txt",
        .{ .mode = .read_only },
    );
    defer file.close();

    //? Get file size
    const fileStats = try file.stat();
    const fileSize = fileStats.size;
    try writer.print("File Size: {d} \n", .{fileSize});

    //? Read Data into buffer
    var reader = file.reader();

    const data: []u8 = try reader.readAllAlloc(allocator, fileSize * 2);
    defer allocator.free(data);

    //? Use data
    var lines = std.mem.split(u8, data, "\n");


    var firstList = ArrayList(LType).init(allocator);
    defer firstList.deinit();

    var secondList = ArrayList(LType).init(allocator);
    defer secondList.deinit();


    while (lines.next()) |uncleanLine| {
        const line = std.mem.trim(u8, uncleanLine, "\r");

        var nums = std.mem.split(u8, line, "   ");

        const firstValue = nums.next() orelse unreachable;
        const secondValue = nums.next() orelse unreachable;

        const num1 = try std.fmt.parseInt(LType, firstValue, 10);
        const num2 = try std.fmt.parseInt(LType, secondValue, 10);


        try firstList.append(num1);
        try secondList.append(num2);
    }


    const firstArray = try firstList.toOwnedSlice();
    defer allocator.free(firstArray);
    const secondArray = try secondList.toOwnedSlice();
    defer allocator.free(secondArray);

    std.mem.sort(LType, firstArray, {}, comptime std.sort.asc(LType));
    std.mem.sort(LType, secondArray, {}, comptime std.sort.asc(LType));

    var sum: LType = 0;

    for (firstArray, 0..) |firstElement, index| {
        sum = sum + abs(firstElement - secondArray[index]);
    }

    var simSum: LType = 0;

    for (firstArray) |firstElement| {
        var counter: LType = 0;
        for (secondArray) |secondElement| {
            if (firstElement == secondElement) {
                counter = counter + 1;
            }
        }
        simSum = simSum + (counter * firstElement);
    }


    try writer.print("Result 1: {d} \n", .{sum});
    try writer.print("Result 2: {d} \n", .{simSum});

    try writer.print("\nFinished!", .{});
}


pub fn abs(num: LType) LType {
    if (num < 0) {
        return num * -1;
    }
    else return num;
}


pub fn printU64(arr: []LType) !void {
    const writer = std.io.getStdOut().writer();

    for (arr) |x| {
        try writer.print("{d} \n", .{x});
    }
}


pub fn printArrL(list: ArrayList(LType)) !void {
    const writer = std.io.getStdOut().writer();

    for (list.items) |x| {
        try writer.print("{d} \n", .{x});
    }
}
