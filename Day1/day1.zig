const std = @import("std");
const Allocator = std.mem.Allocator;
const expect = std.testing.expect;
const ArrayList = std.ArrayList;

const Utils = @import("Utils.zig");

const sort = std.mem.sort;
const dbg = std.debug;

const LType: type = i64;

pub fn main() !void {
    const writer = std.io.getStdOut().writer();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();


    const data = try Utils.readFileIntoArr(allocator, "Data.txt");
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
        sum = sum + Utils.abs(LType, firstElement - secondArray[index]);
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





