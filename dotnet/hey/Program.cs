using System.Runtime.InteropServices;


Console.WriteLine("content-type: text/plain;utf-8");
Console.WriteLine("");

Console.WriteLine("👋 Hello, World! 🌍");
Console.WriteLine($"OS Architecture: {RuntimeInformation.OSArchitecture}");
Console.WriteLine($"Current time (UTC): {DateTime.UtcNow.ToLongTimeString()}");
