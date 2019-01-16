using System;
using System.Collections.Generic;

namespace BufferOptions
{
    /// <summary>
    /// this class is written to hold functionality
    /// for programs that want to produce output
    /// as soon as an input is processed
    /// </summary>
    public static class DisallowOutputBuffer
    {
        /// <summary>
        /// Returns an the ouput which is the result of running the lambda on an input
        /// </summary>
        public static IEnumerable<string> FilterIterator(string token, Func<string, string> filterLogic)
        {
            string output = filterLogic(token);
            yield return output;
        }
    }
}
