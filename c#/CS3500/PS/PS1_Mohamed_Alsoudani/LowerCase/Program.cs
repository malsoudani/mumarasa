using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LowerCase
{
    class Program
    {
        static void Main(string[] args)
        {
            // making sure that we have exactly one parameter
            if (args.Length != 1 && string.IsNullOrEmpty(args[0]))
            {
                Console.Error.WriteLine("provide one string to lower-case it");
                System.Environment.Exit(1);
            }
            string subject = args[0];
            string[] subjectTokens = subject.Split(new char[0], StringSplitOptions.RemoveEmptyEntries);
        }
    }
}
