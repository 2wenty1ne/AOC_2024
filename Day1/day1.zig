const std = @import("std");
const Allocator = std.mem.Allocator;
const sort = std.mem.sort;

const LType: type = u8; 

pub fn main() !void {
    //const Self = @This();
    const writer = std.io.getStdOut().writer();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    // &[_]LType{ 3, 4, 2, 1, 3, 3}
    //? Get file
    const file = try std.fs.cwd().openFile("testData.txt", .{.mode=.read_only},);
    defer file.close();

    //? Get file size
    const fileStats = try file.stat();
    const fileSize = fileStats.size;
    try writer.print("File Size: {d} \n", .{fileSize});


    //? Read data into buffer
    const buffer = try allocator.alloc(u8, fileSize);
    defer allocator.free(buffer);
    try file.seekTo(0);
    _ = try file.readAll(buffer);

    try writer.print("File Content: \n{s}\n", .{buffer});

    

    // sort(u8, list1, {}, comptime.sort.desc(u8));


    try writer.print("Finished!", .{});
}


// pub fn printArrL(list: ArrayList(LType)) !void {
//     const writer = std.io.getStdOut().writer();

//     for (list.items) |x| {
//         try writer.print("{d} \n",.{x});
//     }
// }