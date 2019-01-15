using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using BufferOptions;

namespace LowerCase
{
    class Program
    {
        /// <summary>
        /// returns the words lowercased without an output buffer "returns it as it processes each of the word of the inputs on its own"
        /// </summary>
        static void Main(string[] args)
        {
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                string[] lineTokens = line.Split(new char[] {' ', '\t', '\r', '\n'}, StringSplitOptions.RemoveEmptyEntries);
                foreach (string token in lineTokens)
                {
                    foreach (string output in DisallowOutputBuffer.FilterIterator(token, s => s.ToLower()))
                    {
                        Console.WriteLine(output);
                    }
                }
            }
        }
    }
}

// notice that through using an iterator as its own class I provide abstraction to the program
// and just inject "business" logic directly to the function as a lambda
// now I can document this as a function that it could be used by another developer
// without having to worry about how the iteration is actually handled
// plus I can use this for as many filters as I want since I'm customizing my logic on the fly
// another benefit is that I can change the entire filter with changing a single function
// try: ToLower() -> ToUpper()